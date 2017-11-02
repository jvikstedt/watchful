package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/pkg/exec"
)

type executableResp struct {
	Identifier string `json:"identifier"`
	exec.Instruction
}

func (h handler) executableAll(w http.ResponseWriter, r *http.Request) {
	executors := h.manager.Executables()

	responses := []executableResp{}
	for key, v := range executors {
		responses = append(responses, executableResp{Identifier: key, Instruction: v.Instruction()})
	}

	json.NewEncoder(w).Encode(responses)
}
