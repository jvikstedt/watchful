package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/model"
)

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	job := &model.Job{}

	if err := decoder.Decode(job); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	if h.checkErr(h.storage.Job().Create(job), w, http.StatusUnprocessableEntity) {
		return
	}

	json.NewEncoder(w).Encode(job)
}
