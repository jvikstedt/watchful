package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/model"
)

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	job := &model.Job{}

	if err := decoder.Decode(job); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.JobCreate(job), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(job)
}

func (h handler) jobGetOne(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "jobID")
	if idStr == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	jobID, err := strconv.Atoi(idStr)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job := model.Job{}
	if h.checkErr(h.model.JobGetOne(jobID, &job), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(job)
}

func (h handler) jobUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "jobID")
	if idStr == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	jobID, err := strconv.Atoi(idStr)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job := model.Job{}
	if h.checkErr(h.model.JobGetOne(jobID, &job), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&job); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.JobUpdate(&job), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(job)
}
