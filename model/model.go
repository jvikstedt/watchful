package model

type db interface {
	Close() error
	EnsureTables() error

	JobCreate(*Job) error
	JobGetOne(int, *Job) error

	TaskCreate(*Task) error
	TaskDelete(*Task) error
	TaskGetOne(int, *Task) error
	TaskAllByJobID(int) ([]*Task, error)
}

type Service struct {
	db db
}

func New(db db) *Service {
	return &Service{
		db: db,
	}
}