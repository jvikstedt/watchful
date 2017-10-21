package storage

import "github.com/jvikstedt/watchful/model"

type Service interface {
	Close() error
	EnsureTables() error
	ProjectAll() ([]*model.Project, error)
	JobCreate(string) (*model.Job, error)
	JobGetOne(int) (*model.Job, error)
	TaskCreate(int, string) (*model.Task, error)
	TaskDelete(int) (*model.Task, error)
	TaskGetOne(int) (*model.Task, error)
	TasksByJobID(int) ([]*model.Task, error)
	InputCreate(int, string) (*model.Input, error)
	InputGetOne(int) (*model.Input, error)
}
