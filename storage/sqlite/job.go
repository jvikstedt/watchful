package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) JobGetOne(id int) (*model.Job, error) {
	job := &model.Job{}

	err := s.db.Get(job, "SELECT jobs.* FROM jobs WHERE id=$1", id)
	return job, err
}

func (s *sqlite) JobCreate(name string) (*model.Job, error) {
	result, err := s.db.Exec(`INSERT INTO jobs (name, created_at, updated_at) VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.JobGetOne(int(id))
}
