package model

import (
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	JobID     int        `json:"jobID" db:"job_id"`
	Executor  string     `json:"executor"`
	Inputs    []*Input   `json:"inputs"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}
