package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/model"
)

type taskResponse struct {
	model.Task
	Inputs []model.Input `json:"inputs"`
}

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

	inputs, err := h.model.InputAllByJobID(jobID)
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	responses := []taskResponse{}
	for _, task := range tasks {
		tr := taskResponse{
			Task:   task,
			Inputs: []model.Input{},
		}

		for _, input := range inputs {
			if input.TaskID == task.ID {
				tr.Inputs = append(tr.Inputs, input)
			}
		}
		responses = append(responses, tr)
	}

	json.NewEncoder(w).Encode(responses)
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

	response := taskResponse{
		Task:   task,
		Inputs: []model.Input{},
	}
	executor := h.manager.Executors()[task.Executor]
	for _, i := range executor.Instruction().Input {
		input := model.Input{
			TaskID: task.ID,
			Name:   i.Name,
			Value:  "",
		}
		if h.checkErr(h.model.InputCreate(&input), w, http.StatusUnprocessableEntity) {
			return
		}
		response.Inputs = append(response.Inputs, input)
	}

	json.NewEncoder(w).Encode(response)
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
