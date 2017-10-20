package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/manager"
)

type executorResp struct {
	Identifier string `json:"identifier"`
	manager.Instruction
}

func (h handler) executorAll(w http.ResponseWriter, r *http.Request) {
	executors := h.manager.Executors()

	responses := []executorResp{}
	for key, v := range executors {
		responses = append(responses, executorResp{Identifier: key, Instruction: v.Instruction()})
	}

	json.NewEncoder(w).Encode(responses)
}
