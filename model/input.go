package model

import (
	"time"
)

type Input struct {
	ID        int        `json:"id"`
	TaskID    int        `json:"taskID" db:"task_id"`
	Name      string     `json:"name"`
	Value     string     `json:"value"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) InputAllByJobID(jobID int) ([]*Input, error) {
	result := []*Input{}
	return result, s.db.Call(func(q Querier) error {
		inputs, err := s.db.InputAllByJobID(q, jobID)
		result = inputs
		return err
	}, false)
}

func (s *Service) InputUpdate(input *Input) error {
	return s.db.Call(func(q Querier) error {
		return s.db.InputUpdate(q, input)
	}, false)
}

func (s *Service) InputCreate(input *Input) error {
	return s.db.Call(func(q Querier) error {
		return s.db.InputCreate(q, input)
	}, false)
}

func (s *Service) InputGetOne(id int, input *Input) error {
	return s.db.Call(func(q Querier) error {
		return s.db.InputGetOne(q, id, input)
	}, false)
}
