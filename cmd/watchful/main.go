package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jvikstedt/watchful/pkg/exec"
	"github.com/jvikstedt/watchful/pkg/exec/builtin"
	"github.com/jvikstedt/watchful/pkg/handler"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/storage/sqlite"
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

	manager := exec.NewManager(logger, modelService)
	manager.RegisterExecutor(builtin.Equal{})
	manager.RegisterExecutor(builtin.HTTP{})
	manager.RegisterExecutor(builtin.JSON{})

	go manager.Run()

	http.Handle("/", handler.New(logger, modelService, manager))
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

	if err := manager.Shutdown(); err != nil {
		logger.Printf("Unable to shutdown manager: %v", err)
	}
}
