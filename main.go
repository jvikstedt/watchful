package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jvikstedt/watchful/builtin/executor"
	"github.com/jvikstedt/watchful/handler"
	"github.com/jvikstedt/watchful/manager"
	"github.com/jvikstedt/watchful/model"
	"github.com/jvikstedt/watchful/storage/sqlite"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	storage, err := sqlite.New("./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	model := model.New(storage)

	manager := manager.NewService(logger, model)
	manager.RegisterExecutor(executor.Equal{})
	manager.RegisterExecutor(executor.HTTP{})
	manager.RegisterExecutor(executor.JSON{})

	go manager.Run()

	http.Handle("/", handler.New(logger, model, manager))
	server := &http.Server{Addr: ":" + port}

	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		<-sigquit

		if err := server.Shutdown(context.Background()); err != nil {
			logger.Printf("Unable to shutdown server: %v", err)
		}
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Printf("%v", err)
	} else {
		logger.Println("Server closed!")
	}

	if err := manager.Shutdown(); err != nil {
		logger.Printf("Unable to shutdown manager: %v", err)
	}
}
