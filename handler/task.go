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

	tasks, err := h.model.TaskAllByJobID(jobID)
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h handler) taskCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	task := model.Task{}

	if err := decoder.Decode(&task); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.TaskCreate(&task), w, http.StatusUnprocessableEntity) {
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

	task := model.Task{ID: taskID}
	if h.checkErr(h.model.TaskDelete(&task), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(task)
}
