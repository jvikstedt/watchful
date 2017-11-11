package model

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful"
)

type Input struct {
	ID           int                `json:"id"`
	TaskID       int                `json:"taskID" db:"task_id"`
	Name         string             `json:"name"`
	Value        *InputValue        `json:"value"`
	Dynamic      bool               `json:"dynamic"`
	SourceTaskID *int               `json:"sourceTaskID" db:"source_task_id"`
	SourceName   string             `json:"sourceName" db:"source_name"`
	Type         watchful.ParamType `json:"type" db:"type"`
	CreatedAt    time.Time          `json:"createdAt" db:"created_at" `
	UpdatedAt    time.Time          `json:"updatedAt" db:"updated_at"`
	DeletedAt    *time.Time         `json:"deletedAt" db:"deleted_at"`
}

func (i *Input) Create(e sqlx.Ext) error {
	result, err := e.Exec(`INSERT INTO
		inputs (task_id, name, value, dynamic, source_task_id, source_name, type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		i.TaskID, i.Name, i.Value, i.Dynamic, i.SourceTaskID, i.SourceName, i.Type)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return InputGetOne(e, int(id), i)
}

func (i *Input) Update(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE inputs SET
		value = ?,
		dynamic = ?,
		source_task_id = ?,
		source_name = ?,
		type = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, i.Value, i.Dynamic, i.SourceTaskID, i.SourceName, i.Type, i.ID)
	if err != nil {
		return err
	}

	return InputGetOne(e, i.ID, i)
}

func InputAllByJobID(e sqlx.Ext, jobID int) ([]*Input, error) {
	inputs := []*Input{}
	err := sqlx.Select(e, &inputs, `SELECT inputs.* FROM jobs
		LEFT JOIN tasks ON tasks.job_id=jobs.id
		INNER JOIN inputs ON inputs.task_id=tasks.id
		WHERE jobs.id = ? AND inputs.deleted_at IS NULL`, jobID)
	return inputs, err
}

func InputGetOne(e sqlx.Ext, id int, input *Input) error {
	return sqlx.Get(e, input, "SELECT inputs.* FROM inputs WHERE id=$1", id)
}
