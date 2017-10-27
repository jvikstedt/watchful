package model

import (
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	JobID     int        `json:"jobID" db:"job_id"`
	Executor  string     `json:"executor"`
	Inputs    []*Input   `json:"inputs"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) TaskCreate(task *Task) error {
	return s.db.Call(func(q Querier) error {
		inputs := task.Inputs
		err := s.db.TaskCreate(q, task)
		if err != nil {
			return err
		}
		if len(inputs) > 0 {
			for _, i := range inputs {
				i.TaskID = task.ID
				err = s.db.InputCreate(q, i)
				if err != nil {
					return err
				}
			}
		}
		return err
	}, true)
}

func (s *Service) TaskUpdate(task *Task) error {
	return s.db.Call(func(q Querier) error {
		return s.db.TaskUpdate(q, task)
	}, false)
}

func (s *Service) TaskDelete(task *Task) error {
	return s.db.Call(func(q Querier) error {
		return s.db.TaskDelete(q, task)
	}, false)
}

func (s *Service) TaskGetOne(id int, task *Task) error {
	return s.db.Call(func(q Querier) error {
		return s.db.TaskGetOne(q, id, task)
	}, false)
}

func (s *Service) TaskAllByJobID(jobID int) ([]*Task, error) {
	result := []*Task{}
	return result, s.db.Call(func(q Querier) error {
		tasks, err := s.db.TaskAllByJobID(q, jobID)
		result = tasks
		return err
	}, false)
}
