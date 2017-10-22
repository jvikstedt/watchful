package storage

import "github.com/jvikstedt/watchful/model"

type Service interface {
	Close() error
	EnsureTables() error
	Job() Job
	Task() Task
}

type Job interface {
	Create(string) (*model.Job, error)
	GetOne(int) (*model.Job, error)
}

type Task interface {
	Create(int, string) (*model.Task, error)
	Delete(int) (*model.Task, error)
	GetOne(int) (*model.Task, error)
	AllByJobID(int) ([]*model.Task, error)
}
