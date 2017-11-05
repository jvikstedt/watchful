package model

import "time"

type ResultStatus string

const (
	ResultStatusWaiting = "waiting"
	ResultStatusSuccess = "success"
	ResultStatusError   = "error"
)

type Result struct {
	ID          int           `json:"id"`
	UUID        string        `json:"uuid" db:"uuid"`
	TestRun     bool          `json:"testRun" db:"test_run"`
	JobID       int           `json:"jobID" db:"job_id"`
	Status      ResultStatus  `json:"status" db:"status"`
	ResultItems []*ResultItem `json:"resultItems"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at" `
	UpdatedAt   time.Time     `json:"updatedAt" db:"updated_at"`
}

func (s *Service) ResultGetOneByUUID(uuid string, result *Result) error {
	err := s.db.ResultGetOneByUUID(uuid, result)
	if err != nil {
		return err
	}

	resultItems, err := s.db.ResultItemAllByResultID(result.ID)
	if err != nil {
		return err
	}

	result.ResultItems = resultItems

	return nil
}
