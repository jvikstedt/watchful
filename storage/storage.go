package storage

import "github.com/jvikstedt/watchful/model"

type Service interface {
	Close() error
	EnsureTables() error
	JobCreate(*model.Job) error
	TaskCreate(*model.Task) error
	TaskDelete(int) error
	TasksByJobID(int, *[]model.Task) error
}
