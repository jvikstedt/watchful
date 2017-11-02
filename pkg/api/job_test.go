package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
		{name: "valid", payload: fmt.Sprintf(`{"name": "memory scanner"}`), expectedStatus: http.StatusCreated, expectedName: "memory scanner"},
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

func TestJobGetOne(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	if err := modelService.DB().JobCreate(&job); err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%d", job.ID), "")
		if statusCode != http.StatusOK {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
		}
	})

	t.Run("not found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%d", 999), "")
		if statusCode != http.StatusNotFound {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusNotFound, statusCode)
		}
	})
}

func TestJobUpdate(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	if err := modelService.DB().JobCreate(&job); err != nil {
		t.Fatal(err)
	}

	tt := []struct {
		name           string
		id             int
		payload        string
		expectedStatus int
		expectedName   string
	}{
		{name: "not found", id: 999, payload: "", expectedStatus: http.StatusNotFound},
		{name: "empty body", id: job.ID, payload: "", expectedStatus: http.StatusUnprocessableEntity},
		{name: "no name", id: job.ID, payload: `{"name": ""}`, expectedStatus: http.StatusUnprocessableEntity},
		{name: "valid name", id: job.ID, payload: `{"name": "something"}`, expectedStatus: http.StatusOK, expectedName: "something"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, body := makeRequest(t, "PUT", fmt.Sprintf("/jobs/%d", tc.id), tc.payload)

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

func TestJobTestRun(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	err := modelService.DB().JobCreate(&job)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("ok", func(t *testing.T) {
		statusCode, body := makeRequest(t, "POST", fmt.Sprintf("/jobs/%d/test_run", job.ID), "")
		if statusCode != http.StatusOK {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
		}
		if strings.TrimSpace(body.String()) != `"test"` {
			t.Fatalf("Expected response to be %s but got %s", `"test"`, body.String())
		}
	})

	t.Run("not found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "POST", fmt.Sprintf("/jobs/%d/test_run", 999), "")
		if statusCode != http.StatusNotFound {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusNotFound, statusCode)
		}
	})
}
