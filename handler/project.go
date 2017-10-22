package handler

import (
	"encoding/json"
	"net/http"
)

func (h handler) projectAll(w http.ResponseWriter, r *http.Request) {

	projects, err := h.storage.ProjectAll()
	if h.checkErr(err, w, http.StatusInternalServerError) {
		return
	}

	json.NewEncoder(w).Encode(projects)
}
