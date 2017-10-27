package sqlite

import (
	"github.com/jvikstedt/watchful/model"
)

func (s *sqlite) JobGetOne(q model.Querier, id int, job *model.Job) error {
	return q.QueryRow("SELECT jobs.* FROM jobs WHERE id=$1", id).Scan(&job.ID, &job.Name, &job.Active, &job.CreatedAt, &job.UpdatedAt)
}

func (s *sqlite) JobCreate(q model.Querier, job *model.Job) error {
	result, err := q.Exec(`INSERT INTO jobs (name, active, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, job.Name, job.Active)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.JobGetOne(q, int(id), job)
}

func (s *sqlite) JobUpdate(q model.Querier, job *model.Job) error {
	_, err := q.Exec(`UPDATE jobs SET name = ?, active = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, job.Name, job.Active, job.ID)
	if err != nil {
		return err
	}

	return s.JobGetOne(q, job.ID, job)
}
