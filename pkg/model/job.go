package model

import (
	"fmt"
	"time"
)

type Job struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (s *Service) ValidateJob(job *Job) error {
	if len(job.Name) == 0 {
		return fmt.Errorf("Job name length can't be 0")
	}
	return nil
}
