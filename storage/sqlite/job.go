package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) JobGetOne(id int, job *model.Job) error {
	return s.q.Get(job, "SELECT jobs.* FROM jobs WHERE id=$1", id)
}

func (s *sqlite) JobCreate(job *model.Job) error {
	result, err := s.q.Exec(`INSERT INTO jobs (name, active, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, job.Name, job.Active)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.JobGetOne(int(id), job)
}

func (s *sqlite) JobUpdate(job *model.Job) error {
	_, err := s.q.Exec(`UPDATE jobs SET name = ?, active = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, job.Name, job.Active, job.ID)
	if err != nil {
		return err
	}

	return s.JobGetOne(job.ID, job)
}
