package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"

	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/api"
	"github.com/jvikstedt/watchful/pkg/exec"
	"github.com/jvikstedt/watchful/pkg/exec/builtin"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/schedule"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	rootDir, err := getRootDir()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(filepath.Join(rootDir, "watchful.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)

	db, err := model.NewDB("sqlite3", filepath.Join(rootDir, "watchful.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := model.EnsureTables(db); err != nil {
		log.Fatal(err)
	}

	jobs, err := model.JobAll(db)
	if err != nil {
		log.Fatal(err)
	}

	execService := exec.New(logger, db)
	execService.RegisterExecutable(builtin.Equal{})
	execService.RegisterExecutable(builtin.HTTP{})
	execService.RegisterExecutable(builtin.DiskInfo{})
	execService.RegisterExecutable(builtin.MoreThan{})
	// execService.RegisterExecutable(builtin.JSON{})

	extFolder := filepath.Join(rootDir, "ext")
	log.Printf("Loading extension from folder: %s\n", extFolder)
	err = exec.SearchExtensions(extFolder, func(e watchful.Executable, err error) {
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Registering extension: %s\n", e.Identifier())
		execService.RegisterExecutable(e)
	})
	if err != nil {
		log.Println(err)
	}

	go execService.Run()
	defer execService.Shutdown()

	// Scheduler
	scheduler := schedule.NewCronScheduler(logger)
	go scheduler.Start()
	defer scheduler.Stop()

	for _, job := range jobs {
		if job.Active {
			scheduler.AddEntry(schedule.EntryID(job.ID), job.Cron, func(id schedule.EntryID) {
				execService.AddJob(job, false)
			})
		}
	}

	http.Handle("/", api.New(logger, db, execService, scheduler))
	server := &http.Server{Addr: ":" + port}

	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		<-sigquit

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Printf("Unable to shutdown server: %v", err)
		}
	}()

	logger.Printf("Server starting on :%s\n", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Printf("%v", err)
	} else {
		logger.Println("Server closed!")
	}
}

func getRootDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	homeDir := usr.HomeDir

	rootDir := filepath.Join(homeDir, ".watchful")
	extPath := filepath.Join(rootDir, "ext")

	os.MkdirAll(extPath, os.ModePerm)

	return rootDir, nil
}
