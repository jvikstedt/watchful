package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jvikstedt/watchful/pkg/exec"
	"github.com/jvikstedt/watchful/pkg/model"
)

var EmptyObject = struct{}{}

func New(logger *log.Logger, model *model.Service, exec *exec.Service) http.Handler {
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

	h := handler{logger, model, exec}

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/jobs", func(r chi.Router) {
			r.Post("/", h.jsonResponseHandler(h.jobCreate))
			r.Route("/{jobID}", func(r chi.Router) {
				r.Get("/tasks", h.taskAll)
				r.Get("/", h.jobGetOne)
				r.Put("/", h.jobUpdate)
				r.Post("/test_run", h.jobTestRun)
			})
		})

		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.taskCreate)
			r.Route("/{taskID}", func(r chi.Router) {
				r.Delete("/", h.taskDelete)
				r.Put("/", h.taskUpdate)
			})
		})

		r.Route("/results", func(r chi.Router) {
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", h.resultGetOne)
			})
		})

		r.Route("/inputs", func(r chi.Router) {
			r.Route("/{inputID}", func(r chi.Router) {
				r.Put("/", h.inputUpdate)
			})
		})

		r.Route("/executors", func(r chi.Router) {
			r.Get("/", h.executorAll)
		})
	})
	return r
}

type handler struct {
	log   *log.Logger
	model *model.Service
	exec  *exec.Service
}

func (h handler) checkErr(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(statusCode)
		return true
	}
	return false
}

func (h handler) getURLParamInt(r *http.Request, key string) (int, error) {
	idStr := chi.URLParam(r, key)
	if idStr == "" {
		return 0, fmt.Errorf("URL param %s was empty", key)
	}
	return strconv.Atoi(idStr)
}

func (h handler) jsonResponseHandler(handleFunc func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := handleFunc(w, r)
		if err != nil {
			h.log.Println(err)
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			h.log.Printf("Could not encode response to output: %v", err)
		}
	}
}
