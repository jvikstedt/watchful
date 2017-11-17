package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful/pkg/model"
)

func TestTaskGetAll(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	task := &model.Task{JobID: job.ID}
	if err := task.Create(db); err != nil {
		t.Fatal(err)
	}

	statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%d/tasks", task.JobID), "")
	if statusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
	}
}

func TestTaskCreate(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		name               string
		payload            string
		expectedStatus     int
		expectedExecutable string
	}{
		{name: "valid", payload: fmt.Sprintf(`{"jobID": %d, "executable": "http"}`, job.ID), expectedStatus: http.StatusCreated, expectedExecutable: "http"},
		{name: "empty payload", payload: "", expectedStatus: http.StatusUnprocessableEntity},
		{name: "invalid executable", payload: fmt.Sprintf(`{"jobID": %d, "executable": "invalid"}`, job.ID), expectedStatus: http.StatusNotFound},
		{name: "invalid jobID", payload: fmt.Sprintf(`{"jobID": %d, "executable": "http"}`, 999), expectedStatus: http.StatusUnprocessableEntity},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, body := makeRequest(t, "POST", "/tasks", tc.payload)

			if statusCode != tc.expectedStatus {
				t.Fatalf("Expected StatusCode %d but got %d", tc.expectedStatus, statusCode)
			}

			task := model.Task{}
			json.NewDecoder(body).Decode(&task)

			if task.Executable != tc.expectedExecutable {
				t.Fatalf("Expected Executable %s but got %s", tc.expectedExecutable, task.Executable)
			}
		})
	}
}

func TestTaskUpdate(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	task := &model.Task{JobID: job.ID, Executable: "http"}
	if err := task.Create(db); err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		name               string
		id                 int
		payload            string
		expectedStatus     int
		expectedExecutable string
	}{
		{name: "not found", id: 999, payload: "", expectedStatus: http.StatusNotFound},
		{name: "empty body", id: task.ID, payload: "", expectedStatus: http.StatusUnprocessableEntity},
		{name: "valid name", id: task.ID, payload: `{"executable": "equal"}`, expectedStatus: http.StatusOK, expectedExecutable: "equal"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, body := makeRequest(t, "PUT", fmt.Sprintf("/tasks/%d", tc.id), tc.payload)

			if statusCode != tc.expectedStatus {
				t.Fatalf("Expected StatusCode %d but got %d", tc.expectedStatus, statusCode)
			}

			task := model.Task{}
			json.NewDecoder(body).Decode(&task)

			if task.Executable != tc.expectedExecutable {
				t.Fatalf("Expected Executable %s but got %s", tc.expectedExecutable, task.Executable)
			}
		})
	}
}

func TestTaskDelete(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	task := &model.Task{JobID: job.ID}
	if err := task.Create(db); err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "DELETE", fmt.Sprintf("/tasks/%d", task.ID), "")
		if statusCode != http.StatusOK {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
		}
	})

	t.Run("not found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "DELETE", fmt.Sprintf("/tasks/%d", 999), "")
		if statusCode != http.StatusNotFound {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusNotFound, statusCode)
		}
	})
}
