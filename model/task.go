package model

import (
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	JobID     int        `json:"jobID" db:"job_id"`
	Executor  string     `json:"executor"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) TaskCreate(task *Task) error {
	return s.db.TaskCreate(task)
}

func (s *Service) TaskDelete(task *Task) error {
	return s.db.TaskDelete(task)
}

func (s *Service) TaskGetOne(id int, task *Task) error {
	return s.db.TaskGetOne(id, task)
}

func (s *Service) TaskAllByJobID(id int) ([]*Task, error) {
	return s.db.TaskAllByJobID(id)
}
