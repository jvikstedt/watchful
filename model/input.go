package model

import (
	"time"
)

type Input struct {
	ID        int        `json:"id"`
	TaskID    int        `json:"taskID" db:"task_id"`
	Name      string     `json:"name"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) InputAllByJobID(id int) ([]Input, error) {
	return s.db.InputAllByJobID(id)
}
