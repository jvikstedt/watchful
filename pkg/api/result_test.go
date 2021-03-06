package api_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful/pkg/model"
)

func TestResultGetOne(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	result := &model.Result{UUID: "abc", JobID: job.ID}
	if err := result.Create(db); err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/results/%s", result.UUID), "")
		if statusCode != http.StatusOK {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
		}
	})

	t.Run("not found", func(t *testing.T) {
		statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%s", "999"), "")
		if statusCode != http.StatusNotFound {
			t.Fatalf("Expected StatusCode %d but got %d", http.StatusNotFound, statusCode)
		}
	})
}

func TestResultGetAll(t *testing.T) {
	job := &model.Job{Name: "memory scanner"}
	if err := job.Create(db); err != nil {
		t.Fatal(err)
	}

	result := &model.Result{JobID: job.ID}
	if err := result.Create(db); err != nil {
		t.Fatal(err)
	}

	statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%d/results", result.JobID), "")
	if statusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
	}
}
