package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

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

func (r *Result) Create(e sqlx.Ext) error {
	result, err := e.Exec(`INSERT INTO
		results (uuid, test_run, job_id, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, r.UUID, r.TestRun, r.JobID, r.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return ResultGetOne(e, int(id), r)
}

func (r *Result) Update(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE results SET
		uuid = ?,
		test_run = ?,
		job_id = ?,
		status = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, r.UUID, r.TestRun, r.JobID, r.Status, r.ID)
	if err != nil {
		return err
	}

	return ResultGetOne(e, r.ID, r)
}

func ResultGetOneByUUIDWithItems(e sqlx.Ext, uuid string, result *Result) error {
	err := ResultGetOneByUUID(e, uuid, result)
	if err != nil {
		return err
	}

	resultItems, err := ResultItemAllByResultID(e, result.ID)
	if err != nil {
		return err
	}

	result.ResultItems = resultItems

	return nil
}

func ResultGetOne(e sqlx.Ext, id int, result *Result) error {
	return sqlx.Get(e, result, "SELECT * FROM results WHERE id=$1", id)
}

func ResultGetOneByUUID(e sqlx.Ext, uuid string, result *Result) error {
	return sqlx.Get(e, result, "SELECT * FROM results WHERE uuid=$1", uuid)
}

func ResultAllByJobID(e sqlx.Ext, jobID int, limit int, offset int) ([]*Result, error) {
	results := []*Result{}
	err := sqlx.Select(e, &results, `SELECT * FROM results
		WHERE job_id = ?
		ORDER BY created_at
		DESC LIMIT ?
		OFFSET ?`, jobID, limit, offset)
	return results, err
}
