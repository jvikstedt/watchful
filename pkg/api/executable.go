package api

import (
	"net/http"

	"github.com/jvikstedt/watchful"
)

type executableResp struct {
	Identifier string `json:"identifier"`
	watchful.Instruction
}

func (h handler) executableAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	executors := h.exec.Executables()

	responses := []executableResp{}
	for key, v := range executors {
		responses = append(responses, executableResp{Identifier: key, Instruction: v.Instruction()})
	}

	return responses, http.StatusOK, nil
}
