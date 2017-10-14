package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/manager"
)

type executorResp struct {
	Name string `json:"name"`
	manager.Instruction
}

func (h handler) executorAll(w http.ResponseWriter, r *http.Request) {
	executors := h.manager.Executors()

	responses := []executorResp{}
	for key, v := range executors {
		responses = append(responses, executorResp{Name: key, Instruction: v.Instructions()})
	}

	json.NewEncoder(w).Encode(responses)
}
