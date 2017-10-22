package sqlite

import "github.com/jvikstedt/watchful/model"

type sqliteTask struct {
	*sqlite
}

func (s *sqliteTask) Create(task *model.Task) error {
	result, err := s.db.Exec(`INSERT INTO tasks (job_id, executor, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, task.JobID, task.Executor)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.GetOne(int(id), task)
}

func (s *sqliteTask) GetOne(id int, task *model.Task) error {
	return s.db.Get(task, "SELECT * FROM tasks WHERE id=$1", id)
}

func (s *sqliteTask) Delete(task *model.Task) error {
	_, err := s.db.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, task.ID)
	if err != nil {
		return err
	}
	return s.GetOne(task.ID, task)
}

func (s *sqliteTask) AllByJobID(jobID int) ([]*model.Task, error) {
	tasks := []*model.Task{}
	err := s.db.Select(&tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	return tasks, err
}
