package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/model"
)

func (h handler) inputUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "inputID")
	if idStr == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	inputID, err := strconv.Atoi(idStr)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	input := model.Input{}
	if h.checkErr(h.model.InputGetOne(inputID, &input), w, http.StatusNotFound) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.model.InputUpdate(&input), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(input)
}
