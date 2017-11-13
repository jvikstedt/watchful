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

	tasks, err := model.TasksWithInputsByJobID(h.db, jobID)
	if err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	return tasks, http.StatusOK, nil
}

func (h handler) taskCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)

	task := &model.Task{}

	if err := decoder.Decode(task); err != nil {
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

	tx, err := h.db.Beginx()
	if err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	if err := task.Create(tx); err != nil {
		tx.Rollback()
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := tx.Commit(); err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	return task, http.StatusCreated, nil
}

func (h handler) taskUpdate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	task := &model.Task{}
	if err := model.TaskGetOne(h.db, taskID, task); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(task); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := task.Update(h.db); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return task, http.StatusOK, nil
}

func (h handler) taskDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	taskID, err := h.getURLParamInt(r, "taskID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	task := &model.Task{ID: taskID}
	if err := task.Delete(h.db); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	return task, http.StatusOK, nil
}

func (h handler) taskSwapSeq(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)

	ids := struct {
		ID1 int
		ID2 int
	}{}

	if err := decoder.Decode(&ids); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	tx, err := h.db.Beginx()
	if err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	task1 := &model.Task{}
	if err := model.TaskGetOne(tx, ids.ID1, task1); err != nil {
		tx.Rollback()
		return EmptyObject, http.StatusNotFound, err
	}

	task2 := &model.Task{}
	if err := model.TaskGetOne(tx, ids.ID2, task2); err != nil {
		tx.Rollback()
		return EmptyObject, http.StatusNotFound, err
	}

	if err := model.TaskSwapSeq(tx, task1, task2); err != nil {
		tx.Rollback()
		return EmptyObject, http.StatusInternalServerError, err
	}

	if err := tx.Commit(); err != nil {
		return EmptyObject, http.StatusInternalServerError, err
	}

	return EmptyObject, http.StatusOK, nil
}
