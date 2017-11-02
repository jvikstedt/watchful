package handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful"
)

func TestExecutableGetAll(t *testing.T) {
	statusCode, body := makeRequest(t, "GET", "/executables", "")
	if statusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
	}

	executables := []struct {
		Identifier string `json:"identifier"`
		watchful.Instruction
	}{}

	json.NewDecoder(body).Decode(&executables)

	if len(executables) != 2 {
		t.Fatalf("Expected length of executables to be %d but got %d", 2, len(executables))
	}
}
