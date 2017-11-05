package model

import (
	"time"

	"github.com/jvikstedt/watchful"
)

type Input struct {
	ID           int                `json:"id"`
	TaskID       int                `json:"taskID" db:"task_id"`
	Name         string             `json:"name"`
	Value        string             `json:"value"`
	Dynamic      bool               `json:"dynamic"`
	SourceTaskID *int               `json:"sourceTaskID" db:"source_task_id"`
	SourceName   string             `json:"sourceName" db:"source_name"`
	Type         watchful.ParamType `json:"type" db:"type"`
	CreatedAt    time.Time          `json:"createdAt" db:"created_at" `
	UpdatedAt    time.Time          `json:"updatedAt" db:"updated_at"`
	DeletedAt    *time.Time         `json:"deletedAt" db:"deleted_at"`
}
