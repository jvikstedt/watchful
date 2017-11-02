package api_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jvikstedt/watchful/pkg/model"
)

func TestTaskGetAll(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	if err := modelService.DB().JobCreate(&job); err != nil {
		t.Fatal(err)
	}

	task := model.Task{JobID: job.ID}
	if err := modelService.DB().TaskCreate(&task); err != nil {
		t.Fatal(err)
	}

	statusCode, _ := makeRequest(t, "GET", fmt.Sprintf("/jobs/%d/tasks", task.JobID), "")
	if statusCode != http.StatusOK {
		t.Fatalf("Expected StatusCode %d but got %d", http.StatusOK, statusCode)
	}
}

func TestTaskDelete(t *testing.T) {
	job := model.Job{Name: "memory scanner"}
	if err := modelService.DB().JobCreate(&job); err != nil {
		t.Fatal(err)
	}

	task := model.Task{JobID: job.ID}
	if err := modelService.DB().TaskCreate(&task); err != nil {
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
