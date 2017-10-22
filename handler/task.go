package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/model"
)

func (h handler) taskAll(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "jobID")
	if idStr == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	jobID, err := strconv.Atoi(idStr)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	tasks := []model.Task{}
	if h.checkErr(h.storage.TasksByJobID(jobID, &tasks), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h handler) taskCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	task := &model.Task{}

	if err := decoder.Decode(task); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.storage.TaskCreate(task), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h handler) taskDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "taskID")
	if idStr == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	taskID, err := strconv.Atoi(idStr)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.storage.TaskDelete(taskID), w, http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusOK)
}
