package api

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
	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/schedule"
)

type executor interface {
	AddJob(job *model.Job, isTestRun bool) string
	Executables() map[string]watchful.Executable
}

var EmptyObject = struct{}{}

func New(logger *log.Logger, db *sqlx.DB, exec executor, scheduler schedule.Scheduler) http.Handler {
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

	h := handler{logger, db, exec, scheduler}

	r.Route("/", func(r chi.Router) {
		r.Route("/jobs", func(r chi.Router) {
			r.Get("/", h.jsonResponseHandler(h.jobGetAll))
			r.Post("/", h.jsonResponseHandler(h.jobCreate))
			r.Route("/{jobID}", func(r chi.Router) {
				r.Get("/tasks", h.jsonResponseHandler(h.taskAll))
				r.Get("/results", h.jsonResponseHandler(h.resultAll))
				r.Get("/", h.jsonResponseHandler(h.jobGetOne))
				r.Put("/", h.jsonResponseHandler(h.jobUpdate))
				r.Post("/test_run", h.jsonResponseHandler(h.jobTestRun))
			})
		})

		r.Route("/tasks", func(r chi.Router) {
			r.Post("/", h.jsonResponseHandler(h.taskCreate))
			r.Post("/swap_seq", h.jsonResponseHandler(h.taskSwapSeq))
			r.Route("/{taskID}", func(r chi.Router) {
				r.Delete("/", h.jsonResponseHandler(h.taskDelete))
				r.Put("/", h.jsonResponseHandler(h.taskUpdate))
			})
		})

		r.Route("/results", func(r chi.Router) {
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", h.jsonResponseHandler(h.resultGetOne))
			})
		})

		r.Route("/inputs", func(r chi.Router) {
			r.Post("/", h.jsonResponseHandler(h.inputCreate))
			r.Route("/{inputID}", func(r chi.Router) {
				r.Put("/", h.jsonResponseHandler(h.inputUpdate))
				r.Delete("/", h.jsonResponseHandler(h.inputDelete))
			})
		})

		r.Route("/executables", func(r chi.Router) {
			r.Get("/", h.jsonResponseHandler(h.executableAll))
		})
	})
	return r
}

type handler struct {
	log       *log.Logger
	db        *sqlx.DB
	exec      executor
	scheduler schedule.Scheduler
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
