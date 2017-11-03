package model

import (
	"time"
)

type Task struct {
	ID         int        `json:"id"`
	JobID      int        `json:"jobID" db:"job_id"`
	Executable string     `json:"executable"`
	Inputs     []*Input   `json:"inputs"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt  time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt  *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) TaskCreate(task *Task) error {
	return s.db.BeginTx(func(s Storager) error {
		err := s.TaskCreate(task)
		if err != nil {
			return err
		}

		if len(task.Inputs) > 0 {
			for _, i := range task.Inputs {
				i.TaskID = task.ID
				err = s.InputCreate(i)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (s *Service) TasksWithInputsByJobID(jobID int) ([]*Task, error) {
	tasks, err := s.db.TaskAllByJobID(jobID)
	if err != nil {
		return tasks, err
	}

	inputs, err := s.InputAllByJobID(jobID)
	if err != nil {
		return tasks, err
	}

	for _, task := range tasks {
		for _, input := range inputs {
			if input.TaskID == task.ID {
				task.Inputs = append(task.Inputs, input)
			}
		}
	}

	return tasks, nil
}
