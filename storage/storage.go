package storage

import "github.com/jvikstedt/watchful/model"

type Service interface {
	Close() error
	EnsureTables() error
	ProjectAll() ([]*model.Project, error)
	JobCreate(string, []byte) (*model.Job, error)
	JobGetOne(id int) (*model.Job, error)
}
