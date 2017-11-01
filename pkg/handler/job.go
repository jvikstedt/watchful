package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)
	job := &model.Job{}

	err := decoder.Decode(job)
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	err = h.model.ValidateJob(job)
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	err = h.model.DB().JobCreate(job)
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return job, http.StatusCreated, nil
}

func (h handler) jobGetOne(w http.ResponseWriter, r *http.Request) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job := model.Job{}
	if h.checkErr(h.model.DB().JobGetOne(jobID, &job), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(job)
}

func (h handler) jobUpdate(w http.ResponseWriter, r *http.Request) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job := model.Job{}
	if h.checkErr(h.model.DB().JobGetOne(jobID, &job), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&job); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.DB().JobUpdate(&job), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(job)
}

func (h handler) jobTestRun(w http.ResponseWriter, r *http.Request) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job := model.Job{}
	if h.checkErr(h.model.DB().JobGetOne(jobID, &job), w, http.StatusInternalServerError) {
		return
	}

	id := h.manager.AddScheduledJob(&job, true)

	json.NewEncoder(w).Encode(id)
}
