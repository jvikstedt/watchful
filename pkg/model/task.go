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
	CreatedAt  time.Time  `json:"createdAt" db:"created_at" `
	UpdatedAt  time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt  *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (t *Task) Create(db *sqlx.DB) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	result, err := tx.Exec(`INSERT INTO
		tasks (job_id, executable, created_at, updated_at)
		VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, t.JobID, t.Executable)
	if err != nil {
		tx.Rollback()
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := TaskGetOne(tx, int(id), t); err != nil {
		tx.Rollback()
		return err
	}

	if len(t.Inputs) > 0 {
		for _, i := range t.Inputs {
			i.TaskID = t.ID
			err = i.Create(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

func (t *Task) Update(e sqlx.Ext) error {
	_, err := e.Exec(`UPDATE tasks
		SET executable = ?,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = ?`, t.Executable, t.ID)
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
	err := sqlx.Select(e, &tasks, "SELECT * FROM tasks WHERE job_id = ? AND deleted_at IS NULL", jobID)
	return tasks, err
}
