package model

import "log"

type Querier interface {
	JobCreate(*Job) error
	JobUpdate(*Job) error
	JobGetOne(int, *Job) error

	TaskCreate(*Task) error
	TaskUpdate(*Task) error
	TaskDelete(*Task) error
	TaskGetOne(int, *Task) error
	TaskAllByJobID(int) ([]*Task, error)

	InputAllByJobID(int) ([]*Input, error)
	InputCreate(*Input) error
	InputUpdate(*Input) error
	InputGetOne(int, *Input) error

	ResultCreate(*Result) error
	ResultUpdate(*Result) error
	ResultGetOne(int, *Result) error
	ResultGetOneByUUID(string, *Result) error
}

type db interface {
	Querier
	Close() error
	EnsureTables() error
	BeginTx(func(Querier) error) error
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
