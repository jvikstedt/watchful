package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jvikstedt/watchful/manager"
	"github.com/jvikstedt/watchful/model"
)

func New(logger *log.Logger, model *model.Service, manager *manager.Service) http.Handler {
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

	h := handler{logger, model, manager}

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/jobs", func(r chi.Router) {
			r.Post("/", h.jobCreate)
			r.Route("/{jobID}", func(r chi.Router) {
				r.Get("/tasks", h.taskAll)
			})
		})

		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.taskCreate)
			r.Route("/{taskID}", func(r chi.Router) {
				r.Delete("/", h.taskDelete)
			})
		})

		r.Route("/executors", func(r chi.Router) {
			r.Get("/", h.executorAll)
		})
	})
	return r
}

type handler struct {
	log     *log.Logger
	model   *model.Service
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
