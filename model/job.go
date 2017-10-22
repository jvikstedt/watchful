package model

import (
	"time"
)

type Job struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (s *Service) JobCreate(job *Job) error {
	return s.db.JobCreate(job)
}

func (s *Service) JobGetOne(id int, job *Job) error {
	return s.db.JobGetOne(id, job)
}
