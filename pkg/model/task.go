package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Task struct {
	ID         int        `json:"id"`
	JobID      int        `json:"jobID" db:"job_id"`
	Executable string     `json:"executable"`
	Inputs     []*Input   `json:"inputs"`
	Seq        int        `json:"seq"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt  time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt  *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (t *Task) Create(e sqlx.Ext) error {
	var seq int
	if err := sqlx.Get(e, &seq, "SELECT MAX(seq) FROM tasks WHERE job_id = ?", t.JobID); err != nil {
		seq = 0
	}
	t.Seq = seq + 1

	result, err := e.Exec(`INSERT INTO
		tasks (job_id, executable, seq, created_at, updated_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, t.JobID, t.Executable, t.Seq)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if err := TaskGetOne(e, int(id), t); err != nil {
		return err
	}

	if len(t.Inputs) > 0 {
		for _, i := range t.Inputs {
			i.TaskID = t.ID
			err = i.Create(e)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Task) Update(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE tasks SET
		executable = ?,
		seq = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, t.Executable, t.Seq, t.ID)
	if err != nil {
		return err
	}

	return TaskGetOne(e, t.ID, t)
}

func (t *Task) Delete(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?`, t.ID)
	if err != nil {
		return err
	}
	return TaskGetOne(e, t.ID, t)
}

func TaskSwapSeq(e sqlx.Ext, t1 *Task, t2 *Task) error {
	_, err := e.Exec(`UPDATE tasks SET seq = NULL WHERE id IN (?, ?)`, t1.ID, t2.ID)
	if err != nil {
		return err
	}
	_, err = e.Exec(`UPDATE tasks SET seq = CASE id WHEN ? THEN ? WHEN ? THEN ? END WHERE id IN (?, ?)`, t1.ID, t2.Seq, t2.ID, t1.Seq, t1.ID, t2.ID)
	return err
}

func TasksWithInputsByJobID(e sqlx.Ext, jobID int) ([]*Task, error) {
	tasks, err := TaskAllByJobID(e, jobID)
	if err != nil {
		return tasks, err
	}

	inputs, err := InputAllByJobID(e, jobID)
	if err != nil {
		return tasks, err
	}

	for _, task := range tasks {
		for _, input := range inputs {
			if input.TaskID == task.ID {
				task.Inputs = append(task.Inputs, input)
			}
		}
	}

	return tasks, nil
}

func TaskGetOne(e sqlx.Ext, id int, task *Task) error {
	return sqlx.Get(e, task, "SELECT * FROM tasks WHERE id=$1", id)
}

func TaskAllByJobID(e sqlx.Ext, jobID int) ([]*Task, error) {
	tasks := []*Task{}
	err := sqlx.Select(e, &tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL ORDER BY seq", jobID)
	return tasks, err
}
