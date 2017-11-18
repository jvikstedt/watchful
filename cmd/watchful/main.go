package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/api"
	"github.com/jvikstedt/watchful/pkg/exec"
	"github.com/jvikstedt/watchful/pkg/exec/builtin"
	"github.com/jvikstedt/watchful/pkg/model"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	extFolder := os.Getenv("EXT_FOLDER_PATH")
	if extFolder == "" {
		extFolder = "pkg/exec/example"
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	db, err := model.NewDB("sqlite3", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := model.EnsureTables(db); err != nil {
		log.Fatal(err)
	}

	execService := exec.New(logger, db)
	execService.RegisterExecutable(builtin.Equal{})
	execService.RegisterExecutable(builtin.HTTP{})
	// execService.RegisterExecutable(builtin.JSON{})

	if len(extFolder) > 0 {
		log.Printf("Loading extension from folder: %s\n", extFolder)
		err := exec.SearchExtensions(extFolder, func(e watchful.Executable, err error) {
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
	}

	go execService.Run()

	http.Handle("/", api.New(logger, db, execService))
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

	execService.Shutdown()
}
