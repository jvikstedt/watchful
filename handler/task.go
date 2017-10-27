package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/model"
)

func (h handler) taskAll(w http.ResponseWriter, r *http.Request) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	tasks, err := h.model.TaskAllByJobID(jobID)
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	inputs, err := h.model.InputAllByJobID(jobID)
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	for _, task := range tasks {
		for _, input := range inputs {
			if input.TaskID == task.ID {
				task.Inputs = append(task.Inputs, input)
			}
		}
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h handler) taskCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	task := model.Task{}

	if err := decoder.Decode(&task); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	executor := h.manager.Executors()[task.Executor]
	for _, i := range executor.Instruction().Input {
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
	if h.checkErr(h.model.TaskGetOne(taskID, &task), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.TaskUpdate(&task), w, http.StatusUnprocessableEntity) {
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
	if h.checkErr(h.model.TaskDelete(&task), w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(task)
}
