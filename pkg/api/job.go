package api

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/schedule"
)

func (h handler) jobGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobs, err := model.JobAll(h.db)
	if err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}
	return jobs, http.StatusOK, nil
}

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)
	job := &model.Job{}

	if err := decoder.Decode(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := job.Create(h.db); err != nil {
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
	if err := model.JobGetOne(h.db, jobID, &job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	if job.Active {
		h.scheduler.AddEntry(schedule.EntryID(job.ID), job.Cron, func(id schedule.EntryID) {
			h.exec.AddJob(&job, false)
		})
	}
	return job, http.StatusOK, nil
}

func (h handler) jobUpdate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	job := &model.Job{}
	if err := model.JobGetOne(h.db, jobID, job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(job); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := job.Update(h.db); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if job.Active {
		h.scheduler.AddEntry(schedule.EntryID(job.ID), job.Cron, func(id schedule.EntryID) {
			h.exec.AddJob(job, false)
		})
	} else {
		h.scheduler.RemoveEntry(schedule.EntryID(job.ID))
	}

	return job, http.StatusOK, nil
}

func (h handler) jobTestRun(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	job := model.Job{}
	if err := model.JobGetOne(h.db, jobID, &job); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	id := h.exec.AddJob(&job, true)

	return id, http.StatusOK, nil
}
