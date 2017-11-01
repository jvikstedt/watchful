package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) resultGetOne(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	result := model.Result{}
	if h.checkErr(h.model.ResultGetOneByUUID(uuid, &result), w, http.StatusNotFound) {
		return
	}

	json.NewEncoder(w).Encode(result)
}
