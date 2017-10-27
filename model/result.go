package model

import "time"

type ResultStatus string

const (
	ResultStatusWaiting = "waiting"
	ResultStatusDone    = "done"
)

type Result struct {
	ID        int          `json:"id"`
	UUID      string       `json:"uuid" db:"uuid"`
	TestRun   bool         `json:"testRun" db:"test_run"`
	JobID     int          `json:"jobID" db:"job_id"`
	Status    ResultStatus `json:"status" db:"status"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
}
