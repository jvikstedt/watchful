package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) taskAll(w http.ResponseWriter, r *http.Request) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	tasks, err := h.model.TasksWithInputsByJobID(jobID)
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

	executable := h.exec.Executables()[task.Executable]
	for _, i := range executable.Instruction().Input {
		input := model.Input{Name: i.Name}
		task.Inputs = append(task.Inputs, &input)
	}

	if h.checkErr(h.model.TaskCreate(&task), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h handler) taskUpdate(w http.ResponseWriter, r *http.Request) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	task := model.Task{}
	if h.checkErr(h.model.DB().TaskGetOne(taskID, &task), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.DB().TaskUpdate(&task), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h handler) taskDelete(w http.ResponseWriter, r *http.Request) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	task := model.Task{ID: taskID}
	if h.checkErr(h.model.DB().TaskDelete(&task), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(task)
}
