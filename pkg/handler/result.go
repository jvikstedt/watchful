package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jvikstedt/watchful/pkg/model"
)

func (h handler) resultGetOne(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	uuid := chi.URLParam(r, "uuid")

	result := model.Result{}
	if err := h.model.ResultGetOneByUUID(uuid, &result); err != nil {
		return EmptyObject, http.StatusNotFound, err
	}

	return result, http.StatusOK, nil
}
