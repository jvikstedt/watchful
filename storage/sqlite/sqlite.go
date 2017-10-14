package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful/model"
	"github.com/jvikstedt/watchful/storage"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS projects (
	id integer PRIMARY KEY,
	name text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp
);
CREATE TABLE IF NOT EXISTS jobs (
	id integer PRIMARY KEY,
	executor_id text,
	takes_raw text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp
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

func (s *sqlite) JobCreate(executorID string, takesRaw []byte) (*model.Job, error) {
	result, err := s.db.Exec(`INSERT INTO jobs (executor_id, takes_raw, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, executorID, takesRaw)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.JobGetOne(int(id))
}

func (s *sqlite) EnsureTables() error {
	s.db.MustExec(schema)
	return nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
