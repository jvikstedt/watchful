package sqlite

import (
	"github.com/jvikstedt/watchful/model"
)

func (s *sqlite) TaskCreate(q model.Querier, task *model.Task) error {
	result, err := q.Exec(`INSERT INTO tasks (job_id, executor, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, task.JobID, task.Executor)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return s.TaskGetOne(q, int(id), task)
}

func (s *sqlite) TaskGetOne(q model.Querier, id int, task *model.Task) error {
	return q.QueryRow("SELECT * FROM tasks WHERE id=?", id).Scan(&task.ID, &task.JobID, &task.Executor, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt)
}

func (s *sqlite) TaskDelete(q model.Querier, task *model.Task) error {
	_, err := q.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, task.ID)
	if err != nil {
		return err
	}
	return s.TaskGetOne(q, task.ID, task)
}

func (s *sqlite) TaskAllByJobID(q model.Querier, jobID int) ([]*model.Task, error) {
	tasks := []*model.Task{}
	rows, err := q.Query("SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	for rows.Next() {
		t := model.Task{}
		err = rows.Scan(&t.ID, &t.JobID, &t.Executor, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, err
}

func (s *sqlite) TaskUpdate(q model.Querier, task *model.Task) error {
	_, err := q.Exec(`UPDATE tasks SET executor = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, task.Executor, task.ID)
	if err != nil {
		return err
	}

	return s.TaskGetOne(q, task.ID, task)
}
