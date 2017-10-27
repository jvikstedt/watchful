package sqlite

import (
	"github.com/jvikstedt/watchful/model"
)

func (s *sqlite) TaskCreate(task *model.Task) error {
	result, err := s.q.Exec(`INSERT INTO tasks (job_id, executor, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, task.JobID, task.Executor)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.TaskGetOne(int(id), task)
}

func (s *sqlite) TaskGetOne(id int, task *model.Task) error {
	return s.q.Get(task, "SELECT * FROM tasks WHERE id=$1", id)
}

func (s *sqlite) TaskDelete(task *model.Task) error {
	_, err := s.q.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, task.ID)
	if err != nil {
		return err
	}
	return s.TaskGetOne(task.ID, task)
}

func (s *sqlite) TaskAllByJobID(jobID int) ([]*model.Task, error) {
	tasks := []*model.Task{}
	err := s.q.Select(&tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	return tasks, err
}

func (s *sqlite) TaskUpdate(task *model.Task) error {
	_, err := s.q.Exec(`UPDATE tasks SET executor = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, task.Executor, task.ID)
	if err != nil {
		return err
	}

	return s.TaskGetOne(task.ID, task)
}
