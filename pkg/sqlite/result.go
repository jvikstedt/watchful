package sqlite

import (
	"github.com/jvikstedt/watchful/pkg/model"
)

func (s *sqlite) ResultCreate(result *model.Result) error {
	r, err := s.q.Exec(`INSERT INTO results (uuid, test_run, job_id, status, created_at, updated_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, result.UUID, result.TestRun, result.JobID, result.Status)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	return s.ResultGetOne(int(id), result)
}

func (s *sqlite) ResultGetOne(id int, result *model.Result) error {
	return s.q.Get(result, "SELECT * FROM results WHERE id=$1", id)
}

func (s *sqlite) ResultGetOneByUUID(uuid string, result *model.Result) error {
	return s.q.Get(result, "SELECT * FROM results WHERE uuid=$1", uuid)
}

func (s *sqlite) ResultUpdate(result *model.Result) error {
	_, err := s.q.Exec(`UPDATE results SET uuid = ?, test_run = ?, job_id = ?, status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, result.UUID, result.TestRun, result.JobID, result.Status, result.ID)
	if err != nil {
		return err
	}

	return s.ResultGetOne(result.ID, result)
}
