package handler

import (
	"encoding/json"
	"net/http"
)

type checkerResp struct {
	Identifier string `json:"identifier"`
}

func (h handler) checkerAll(w http.ResponseWriter, r *http.Request) {
	checkers := h.manager.Checkers()

	responses := []checkerResp{}
	for key, _ := range checkers {
		responses = append(responses, checkerResp{Identifier: key})
	}

	json.NewEncoder(w).Encode(responses)
}
