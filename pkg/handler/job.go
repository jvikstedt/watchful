package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)
	job := &model.Job{}

	if err := decoder.Decode(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := h.model.ValidateJob(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := h.model.DB().JobCreate(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return job, http.StatusCreated, nil
}

func (h handler) jobGetOne(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	job := model.Job{}
	if err := h.model.DB().JobGetOne(jobID, &job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	return job, http.StatusOK, nil
}

func (h handler) jobUpdate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	job := &model.Job{}
	if err := h.model.DB().JobGetOne(jobID, job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := h.model.ValidateJob(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := h.model.DB().JobUpdate(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return job, http.StatusOK, nil
}

func (h handler) jobTestRun(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	job := model.Job{}
	if err := h.model.DB().JobGetOne(jobID, &job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	id := h.exec.AddScheduledJob(&job, true)

	return id, http.StatusOK, nil
}
