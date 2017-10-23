package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) InputAllByJobID(jobID int) ([]*model.Input, error) {
	inputs := []*model.Input{}
	err := s.db.Select(&inputs, "SELECT inputs.* FROM jobs LEFT JOIN tasks ON tasks.job_id=jobs.id INNER JOIN inputs ON inputs.task_id=tasks.id WHERE jobs.id = ? AND inputs.deleted_at IS NULL", jobID)
	return inputs, err
}

func (s *sqlite) InputGetOne(id int, input *model.Input) error {
	return s.db.Get(input, "SELECT inputs.* FROM inputs WHERE id=$1", id)
}

func (s *sqlite) InputCreate(input *model.Input) error {
	result, err := s.db.Exec(`INSERT INTO inputs (task_id, name, value, created_at, updated_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, input.TaskID, input.Name, input.Value)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.InputGetOne(int(id), input)
}