package model

import (
	"time"
)

type Task struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	JobID     uint       `json:"jobID"`
	Inputs    []Input    `json:"inputs"`
	Executor  string     `json:"executor"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
