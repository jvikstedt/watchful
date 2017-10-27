package sqlite

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful/model"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS jobs (
	id integer PRIMARY KEY,
	name text,
	active integer DEFAULT 0,
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
CREATE TABLE IF NOT EXISTS results (
	id integer PRIMARY KEY,
	uuid text,
	test_run integer DEFAULT 0,
	job_id integer,
	status text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	FOREIGN KEY(job_id) REFERENCES jobs(id)
);
`

type query interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
}

func New(log *log.Logger, filepath string) (*sqlite, error) {
	db, err := sqlx.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	storage := &sqlite{
		log: log,
		db:  db,
		q:   db,
	}

	err = storage.EnsureTables()

	return storage, err
}

type sqlite struct {
	log *log.Logger
	db  *sqlx.DB
	q   query
	tx  *sqlx.Tx
}

func (s *sqlite) BeginTx(callback func(model.Querier) error) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	err = callback(&sqlite{
		log: s.log,
		db:  s.db,
		q:   tx,
		tx:  tx,
	})

	if err != nil {
		if e := tx.Rollback(); e != nil {
			s.log.Println(e)
		}
		return err
	}

	return tx.Commit()
}

func (s *sqlite) EnsureTables() error {
	_, err := s.db.Exec(schema)
	return err
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
