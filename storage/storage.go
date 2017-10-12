package storage

import "github.com/jvikstedt/watchful/model"

type Service interface {
	Close() error
	EnsureTables() error
	ProjectAll() ([]*model.Project, error)
}
