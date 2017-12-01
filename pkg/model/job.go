package model

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Job struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Active    bool          `json:"active"`
	Status    *ResultStatus `json:"status" db:"status"`
	Cron      string        `json:"cron" db:"cron"`
	LastRun   *time.Time    `json:"lastRun" db:"last_run"`
	CreatedAt time.Time     `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time     `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time    `json:"deletedAt" db:"deleted_at"`
}

func (j *Job) Validate() error {
	if len(j.Name) == 0 {
		return fmt.Errorf("Job name length can't be 0")
	}
	return nil
}

func (j *Job) Create(e sqlx.Ext) error {
	if err := j.Validate(); err != nil {
		return err
	}
	result, err := e.Exec(`INSERT INTO
		jobs (name, active, cron, created_at, updated_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, j.Name, j.Active, j.Cron)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return JobGetOne(e, int(id), j)
}

func (j *Job) Update(e sqlx.Ext) error {
	if err := j.Validate(); err != nil {
		return err
	}
	_, err := e.Exec(`UPDATE jobs SET
		name = ?,
		active = ?,
		cron = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, j.Name, j.Active, j.Cron, j.ID)
	if err != nil {
		return err
	}

	return JobGetOne(e, j.ID, j)
}

func JobGetOne(e sqlx.Ext, id int, job *Job) error {
	return sqlx.Get(e, job, "SELECT jobs.* FROM jobs WHERE id=$1", id)
}

func JobAll(e sqlx.Ext) ([]*Job, error) {
	jobs := []*Job{}
	err := sqlx.Select(e, &jobs, `SELECT jobs.*, results.status, results.created_at AS last_run FROM jobs
		LEFT JOIN results ON results.id = (
			SELECT MAX(id) FROM results WHERE job_id = jobs.id
		) WHERE jobs.deleted_at IS NULL`)
	return jobs, err
}
