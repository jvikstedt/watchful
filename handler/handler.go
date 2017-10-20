package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jvikstedt/watchful/manager"
	"github.com/jvikstedt/watchful/storage"
)

func New(logger *log.Logger, storage storage.Service, manager *manager.Service) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
	r.Use(cors.Handler)

	h := handler{logger, storage, manager}

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/projects", func(r chi.Router) {
			r.Get("/", h.projectAll)
		})

		r.Route("/jobs", func(r chi.Router) {
			r.Post("/", h.jobCreate)
		})

		r.Route("/executors", func(r chi.Router) {
			r.Get("/", h.executorAll)
		})

		r.Route("/checkers", func(r chi.Router) {
			r.Get("/", h.checkerAll)
		})
	})
	return r
}

type handler struct {
	log     *log.Logger
	storage storage.Service
	manager *manager.Service
}

func (h handler) checkErr(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(statusCode)
		return true
	}
	return false
}
