package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jvikstedt/watchful/pkg/api"
	"github.com/jvikstedt/watchful/pkg/exec"
	"github.com/jvikstedt/watchful/pkg/exec/builtin"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/sqlite"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	storage, err := sqlite.New(logger, "./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	modelService := model.New(logger, storage)

	execService := exec.New(logger, modelService)
	execService.RegisterExecutable(builtin.Equal{})
	execService.RegisterExecutable(builtin.HTTP{})
	execService.RegisterExecutable(builtin.JSON{})

	go execService.Run()

	http.Handle("/", api.New(logger, modelService, execService))
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
