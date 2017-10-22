package sqlite

import "github.com/jvikstedt/watchful/model"

type sqliteTask struct {
	*sqlite
}

func (db *sqliteTask) Create(jobID int, executor string) (*model.Task, error) {
	result, err := db.Exec(`INSERT INTO tasks (job_id, executor, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, jobID, executor)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return db.GetOne(int(id))
}

func (db *sqliteTask) GetOne(id int) (*model.Task, error) {
	task := &model.Task{}

	err := db.Get(task, "SELECT * FROM tasks WHERE id=$1", id)
	return task, err
}

func (db *sqliteTask) Delete(id int) (*model.Task, error) {
	_, err := db.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	return db.GetOne(id)
}

func (db *sqliteTask) AllByJobID(jobID int) ([]*model.Task, error) {
	tasks := []*model.Task{}
	err := db.Select(&tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	return tasks, err
}
