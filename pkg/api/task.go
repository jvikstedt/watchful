package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) taskAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	jobID, err := h.getURLParamInt(r, "jobID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	tasks, err := h.model.TasksWithInputsByJobID(jobID)
	if err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	return tasks, http.StatusOK, nil
}

func (h handler) taskCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)

	task := model.Task{}

	if err := decoder.Decode(&task); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	executable, ok := h.exec.Executables()[task.Executable]
	if !ok {
		return EmptyObject, http.StatusNotFound, fmt.Errorf("Could not find executable by identifier: %s", task.Executable)
	}
	for _, i := range executable.Instruction().Input {
		input := model.Input{Name: i.Name, Type: i.Type}
		task.Inputs = append(task.Inputs, &input)
	}

	if err := h.model.TaskCreate(&task); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return task, http.StatusCreated, nil
}

func (h handler) taskUpdate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	task := model.Task{}
	if err := h.model.DB().TaskGetOne(taskID, &task); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := h.model.DB().TaskUpdate(&task); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return task, http.StatusOK, nil
}

func (h handler) taskDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	task := model.Task{ID: taskID}
	if err := h.model.DB().TaskDelete(&task); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	return task, http.StatusOK, nil
}
