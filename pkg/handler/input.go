package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) inputUpdate(w http.ResponseWriter, r *http.Request) {
	inputID, err := h.getURLParamInt(r, "inputID")
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	input := model.Input{}
	if h.checkErr(h.model.DB().InputGetOne(inputID, &input), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.DB().InputUpdate(&input), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(input)
}
