package model

import "log"

type Storager interface {
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
	ResultAllByJobID(int, int, int) ([]*Result, error)

	ResultItemCreate(*ResultItem) error
	ResultItemUpdate(*ResultItem) error
	ResultItemGetOne(int, *ResultItem) error
	ResultItemAllByResultID(int) ([]*ResultItem, error)
}

type db interface {
	Storager
	Close() error
	EnsureTables() error
	BeginTx(func(Storager) error) error
}

type Service struct {
	log *log.Logger
	db  db
}

func (s *Service) DB() db {
	return s.db
}

func New(log *log.Logger, db db) *Service {
	return &Service{
		log: log,
		db:  db,
	}
}
