package model

import (
	"time"
)

type Job struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Tasks     []Task     `json:"tasks"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
