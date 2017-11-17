package api

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) inputCreate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	decoder := json.NewDecoder(r.Body)

	input := &model.Input{}

	if err := decoder.Decode(input); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := input.Create(h.db); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return input, http.StatusCreated, nil
}

func (h handler) inputUpdate(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	inputID, err := h.getURLParamInt(r, "inputID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	input := &model.Input{}
	if err := model.InputGetOne(h.db, inputID, input); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(input); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	if err := input.Update(h.db); err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	return input, http.StatusOK, nil
}

func (h handler) inputDelete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	inputID, err := h.getURLParamInt(r, "inputID")
	if err != nil {
		return EmptyObject, http.StatusUnprocessableEntity, err
	}

	input := &model.Input{ID: inputID}
	if err := input.Delete(h.db); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	return input, http.StatusOK, nil
}
