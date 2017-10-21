package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful/model"
	"github.com/jvikstedt/watchful/storage"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS projects (
	id integer PRIMARY KEY,
	name text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS jobs (
	id integer PRIMARY KEY,
	name text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS tasks (
	id integer PRIMARY KEY,
	job_id integer,
	executor text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	FOREIGN KEY(job_id) REFERENCES jobs(id)
);
CREATE TABLE IF NOT EXISTS inputs (
	id integer PRIMARY KEY,
	task_id integer,
	value text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	FOREIGN KEY(task_id) REFERENCES tasks(id)
);
`

func New(filepath string) (storage.Service, error) {
	db, err := sqlx.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	storage := &sqlite{
		db: db,
	}

	storage.EnsureTables()

	return storage, nil
}

type sqlite struct {
	db *sqlx.DB
}

func (s *sqlite) ProjectAll() ([]*model.Project, error) {
	projects := []*model.Project{}
	err := s.db.Select(&projects, "SELECT * FROM projects")
	return projects, err
}

func (s *sqlite) JobGetOne(id int) (*model.Job, error) {
	job := &model.Job{}

	err := s.db.Get(job, "SELECT * FROM jobs WHERE id=$1", id)
	return job, err
}

func (s *sqlite) JobCreate(name string) (*model.Job, error) {
	result, err := s.db.Exec(`INSERT INTO jobs (name, created_at, updated_at) VALUES (?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, name)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.JobGetOne(int(id))
}

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

func (s *sqlite) InputCreate(taskID int, value string) (*model.Input, error) {
	result, err := s.db.Exec(`INSERT INTO inputs (task_id, value, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, taskID, value)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.InputGetOne(int(id))
}

func (s *sqlite) InputGetOne(id int) (*model.Input, error) {
	input := &model.Input{}

	err := s.db.Get(input, "SELECT * FROM inputs WHERE id=$1", id)
	return input, err
}

func (s *sqlite) EnsureTables() error {
	s.db.MustExec(schema)
	return nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
