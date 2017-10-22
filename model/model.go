package model

type db interface {
	Close() error
	EnsureTables() error

	JobCreate(*Job) error
	JobGetOne(int, *Job) error

	TaskCreate(*Task) error
	TaskDelete(*Task) error
	TaskGetOne(int, *Task) error
	TaskAllByJobID(int) ([]Task, error)

	InputAllByJobID(int) ([]Input, error)
	InputCreate(*Input) error
	InputGetOne(int, *Input) error
}

type Service struct {
	db
}

func New(db db) *Service {
	return &Service{
		db: db,
	}
}
