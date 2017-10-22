package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful/storage"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
PRAGMA foreign_keys = ON;
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
	deleted_at timestamp,
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
		DB: db,
	}

	storage.sqliteTask = &sqliteTask{storage}
	storage.sqliteJob = &sqliteJob{storage}

	storage.EnsureTables()

	return storage, nil
}

type sqlite struct {
	*sqlx.DB
	sqliteTask *sqliteTask
	sqliteJob  *sqliteJob
}

func (s *sqlite) EnsureTables() error {
	s.MustExec(schema)
	return nil
}

func (s *sqlite) Job() storage.Job {
	return s.sqliteJob
}

func (s *sqlite) Task() storage.Task {
	return s.sqliteTask
}

func (s *sqlite) Close() error {
	return s.Close()
}
