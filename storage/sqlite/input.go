package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) InputAllByJobID(jobID int) ([]model.Input, error) {
	inputs := []model.Input{}
	err := s.db.Select(&inputs, "SELECT inputs.* FROM jobs LEFT JOIN tasks ON tasks.job_id=jobs.id INNER JOIN inputs ON inputs.task_id=tasks.id WHERE jobs.id = ? AND inputs.deleted_at IS NULL", jobID)
	return inputs, err
}
