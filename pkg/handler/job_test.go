package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful/pkg/model"
)

func TestJobCreate(t *testing.T) {
	tt := []struct {
		name           string
		payload        string
		expectedStatus int
		expectedName   string
	}{
		{name: "memory scanner", payload: fmt.Sprintf(`{"name": "memory scanner"}`), expectedStatus: http.StatusCreated, expectedName: "memory scanner"},
		{name: "empty payload", payload: "", expectedStatus: http.StatusUnprocessableEntity},
		{name: "empty name", payload: fmt.Sprintf(`{"name": ""}`), expectedStatus: http.StatusUnprocessableEntity},
		{name: "no name", payload: fmt.Sprintf(`{}`), expectedStatus: http.StatusUnprocessableEntity},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, body := makeRequest(t, "POST", "/jobs", tc.payload)

			if statusCode != tc.expectedStatus {
				t.Fatalf("Expected StatusCode %d but got %d", tc.expectedStatus, statusCode)
			}

			job := model.Job{}
			json.NewDecoder(body).Decode(&job)

			if job.Name != tc.expectedName {
				t.Fatalf("Expected Name %s but got %s", tc.expectedName, job.Name)
			}
		})
	}
}
