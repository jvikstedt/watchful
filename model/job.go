package model

import (
	"time"
)

type Job struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (s *Service) JobUpdate(job *Job) error {
	return s.db.Call(func(q Querier) error {
		return s.db.JobUpdate(q, job)
	}, false)
}

func (s *Service) JobCreate(job *Job) error {
	return s.db.Call(func(q Querier) error {
		return s.db.JobCreate(q, job)
	}, false)
}

func (s *Service) JobGetOne(id int, job *Job) error {
	return s.db.Call(func(q Querier) error {
		return s.db.JobGetOne(q, id, job)
	}, false)
}
