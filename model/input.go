package model

import (
	"time"
)

type Input struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	TaskID    uint       `json:"taskID"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
