package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) InputAllByJobID(q model.Querier, jobID int) ([]*model.Input, error) {
	inputs := []*model.Input{}
	rows, err := q.Query("SELECT inputs.* FROM jobs LEFT JOIN tasks ON tasks.job_id=jobs.id INNER JOIN inputs ON inputs.task_id=tasks.id WHERE jobs.id = ? AND inputs.deleted_at IS NULL", jobID)
	for rows.Next() {
		t := model.Input{}
		err = rows.Scan(&t.ID, &t.TaskID, &t.Name, &t.Value, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return inputs, err
		}
		inputs = append(inputs, &t)
	}
	return inputs, err
}

func (s *sqlite) InputGetOne(q model.Querier, id int, input *model.Input) error {
	return q.QueryRow("SELECT inputs.* FROM inputs WHERE id=$1", id).Scan(&input.ID, &input.TaskID, &input.Name, &input.Value, &input.CreatedAt, &input.UpdatedAt, &input.DeletedAt)
}

func (s *sqlite) InputCreate(q model.Querier, input *model.Input) error {
	result, err := q.Exec(`INSERT INTO inputs (task_id, name, value, created_at, updated_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, input.TaskID, input.Name, input.Value)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.InputGetOne(q, int(id), input)
}

func (s *sqlite) InputUpdate(q model.Querier, input *model.Input) error {
	_, err := q.Exec(`UPDATE inputs SET value = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, input.Value, input.ID)
	if err != nil {
		return err
	}

	return s.InputGetOne(q, input.ID, input)
}
