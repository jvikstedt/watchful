package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful/pkg/model"
)

func TestInputUpdate(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	if err := modelService.DB().JobCreate(&job); err != nil {
		t.Fatal(err)
	}

	task := model.Task{JobID: job.ID}
	if err := modelService.DB().TaskCreate(&task); err != nil {
		t.Fatal(err)
	}
	input := model.Input{Name: "url", TaskID: task.ID}
	if err := modelService.DB().InputCreate(&input); err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		name           string
		id             int
		payload        string
		expectedStatus int
		expectedValue  string
	}{
		{name: "not found", id: 999, payload: "", expectedStatus: http.StatusNotFound},
		{name: "empty body", id: input.ID, payload: "", expectedStatus: http.StatusUnprocessableEntity},
		{name: "no value", id: input.ID, payload: `{"value": ""}`, expectedStatus: http.StatusOK},
		{name: "valid value", id: input.ID, payload: `{"value": "something"}`, expectedStatus: http.StatusOK, expectedValue: "something"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, body := makeRequest(t, "PUT", fmt.Sprintf("/inputs/%d", tc.id), tc.payload)

			if statusCode != tc.expectedStatus {
				t.Fatalf("Expected StatusCode %d but got %d", tc.expectedStatus, statusCode)
			}

			input := model.Input{}
			json.NewDecoder(body).Decode(&input)

			if input.Value != tc.expectedValue {
				t.Fatalf("Expected Name %s but got %s", tc.expectedValue, input.Value)
			}
		})
	}
}
