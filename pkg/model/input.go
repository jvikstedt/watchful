package model

import (
	"encoding/json"
	"time"

	"github.com/jvikstedt/watchful"
)

type Input struct {
	ID           int                `json:"id"`
	TaskID       int                `json:"taskID" db:"task_id"`
	Name         string             `json:"name"`
	Value        interface{}        `json:"value" db:"-"`
	ValueRaw     []byte             `json:"-" db:"value"`
	Dynamic      bool               `json:"dynamic"`
	SourceTaskID *int               `json:"sourceTaskID" db:"source_task_id"`
	SourceName   string             `json:"sourceName" db:"source_name"`
	Type         watchful.ParamType `json:"type" db:"type"`
	CreatedAt    time.Time          `json:"createdAt" db:"created_at" `
	UpdatedAt    time.Time          `json:"updatedAt" db:"updated_at"`
	DeletedAt    *time.Time         `json:"deletedAt" db:"deleted_at"`
}

func (s *Service) InputUpdate(input *Input) error {
	raw, err := json.Marshal(input.Value)
	if err != nil {
		return err
	}
	input.ValueRaw = raw

	return s.DB().InputUpdate(input)
}

func (s *Service) InputGetOne(id int, input *Input) error {
	err := s.DB().InputGetOne(id, input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(input.ValueRaw, input.Value)
	if err != nil {
		s.log.Println(err)
	}
	return nil
}

func (s *Service) InputAllByJobID(jobID int) ([]*Input, error) {
	inputs, err := s.DB().InputAllByJobID(jobID)
	if err != nil {
		return inputs, err
	}

	for _, i := range inputs {
		err := json.Unmarshal(i.ValueRaw, &i.Value)
		if err != nil {
			s.log.Println(err)
		}
	}

	return inputs, nil
}
