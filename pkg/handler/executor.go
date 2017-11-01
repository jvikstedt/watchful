package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/exec"
)

type executorResp struct {
	Identifier string `json:"identifier"`
	exec.Instruction
}

func (h handler) executorAll(w http.ResponseWriter, r *http.Request) {
	executors := h.exec.Executors()

	responses := []executorResp{}
	for key, v := range executors {
		responses = append(responses, executorResp{Identifier: key, Instruction: v.Instruction()})
	}

	json.NewEncoder(w).Encode(responses)
}
