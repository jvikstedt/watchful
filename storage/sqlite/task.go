package sqlite

import "github.com/jvikstedt/watchful/model"

func (s *sqlite) TaskCreate(jobID int, executor string) (*model.Task, error) {
	result, err := s.db.Exec(`INSERT INTO tasks (job_id, executor, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, jobID, executor)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.TaskGetOne(int(id))
}

func (s *sqlite) TaskGetOne(id int) (*model.Task, error) {
	task := &model.Task{}

	err := s.db.Get(task, "SELECT * FROM tasks WHERE id=$1", id)
	return task, err
}

func (s *sqlite) TaskDelete(id int) (*model.Task, error) {
	_, err := s.db.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	return s.TaskGetOne(id)
}

func (s *sqlite) TasksByJobID(jobID int) ([]*model.Task, error) {
	tasks := []*model.Task{}
	err := s.db.Select(&tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	return tasks, err
}
