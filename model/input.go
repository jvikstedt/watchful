package model

import (
	"time"
)

type Input struct {
	ID        int       `json:"id"`
	TaskID    int       `json:"taskID" db:"task_id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
