package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ResultItem struct {
	ID        int          `json:"id"`
	ResultID  int          `json:"resultID" db:"result_id"`
	TaskID    int          `json:"taskID" db:"task_id"`
	Output    string       `json:"output" db:"output"`
	Error     string       `json:"error"`
	Status    ResultStatus `json:"status"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
}

func (r *ResultItem) Create(e sqlx.Ext) error {
	result, err := e.Exec(`INSERT INTO
		result_items (result_id, task_id, output, error, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, r.ResultID, r.TaskID, r.Output, r.Error, r.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return ResultItemGetOne(e, int(id), r)
}

func (r *ResultItem) Update(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE result_items SET
		result_id = ?,
		task_id = ?,
		output = ?,
		error = ?,
		status = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, r.ResultID, r.TaskID, r.Output, r.Error, r.Status, r.ID)
	if err != nil {
		return err
	}

	return ResultItemGetOne(e, r.ID, r)
}

func ResultItemGetOne(e sqlx.Ext, id int, resultItem *ResultItem) error {
	return sqlx.Get(e, resultItem, "SELECT * FROM result_items WHERE id=$1", id)
}

func ResultItemAllByResultID(e sqlx.Ext, resultID int) ([]*ResultItem, error) {
	resultItems := []*ResultItem{}
	err := sqlx.Select(e, &resultItems, "SELECT * FROM result_items WHERE result_id = ?", resultID)
	return resultItems, err
}
