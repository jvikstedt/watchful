package sqlite

import "github.com/jvikstedt/watchful/model"

type sqliteJob struct {
	*sqlite
}

func (s *sqliteJob) GetOne(id int, job *model.Job) error {
	return s.db.Get(job, "SELECT jobs.* FROM jobs WHERE id=$1", id)
}

func (s *sqliteJob) Create(job *model.Job) error {
	result, err := s.db.Exec(`INSERT INTO jobs (name, created_at, updated_at) VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, job.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.GetOne(int(id), job)
}
