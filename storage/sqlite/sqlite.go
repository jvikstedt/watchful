package sqlite

import (
	"database/sql"
	"log"

	"github.com/jvikstedt/watchful/model"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS jobs (
	id integer PRIMARY KEY,
	name text,
	active integer,
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
	name text,
	value text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	deleted_at timestamp,
	FOREIGN KEY(task_id) REFERENCES tasks(id)
);
`

func New(log *log.Logger, filepath string) (*sqlite, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	storage := &sqlite{
		log: log,
		db:  db,
	}

	storage.EnsureTables()

	return storage, nil
}

type sqlite struct {
	log *log.Logger
	db  *sql.DB
}

func (s *sqlite) EnsureTables() error {
	_, err := s.db.Exec(schema)
	return err
}

func (s *sqlite) Close() error {
	return s.db.Close()
}

func (s *sqlite) Call(callback func(model.Querier) error, transaction bool) error {
	if transaction {
		s.log.Println("Transaction start")
		tx, err := s.db.Begin()
		if err != nil {
			return err
		}

		if err = callback(tx); err != nil {
			tx.Rollback()
			s.log.Printf("Transaction rollback: %v", err)
			return err
		}

		s.log.Println("Transaction commit")
		return tx.Commit()
	}

	return callback(s.db)
}
