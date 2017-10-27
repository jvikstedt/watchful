package model

import (
	"database/sql"
	"log"
)

type Querier interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type db interface {
	Close() error
	EnsureTables() error

	Call(func(Querier) error, bool) error

	JobCreate(Querier, *Job) error
	JobUpdate(Querier, *Job) error
	JobGetOne(Querier, int, *Job) error

	TaskCreate(Querier, *Task) error
	TaskUpdate(Querier, *Task) error
	TaskDelete(Querier, *Task) error
	TaskGetOne(Querier, int, *Task) error
	TaskAllByJobID(Querier, int) ([]*Task, error)

	InputAllByJobID(Querier, int) ([]*Input, error)
	InputCreate(Querier, *Input) error
	InputUpdate(Querier, *Input) error
	InputGetOne(Querier, int, *Input) error
}

type Service struct {
	log *log.Logger
	db
}

func New(log *log.Logger, db db) *Service {
	return &Service{
		log: log,
		db:  db,
	}
}
