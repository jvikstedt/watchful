package api

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

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
