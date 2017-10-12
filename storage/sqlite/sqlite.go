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
);`

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

func (s *sqlite) EnsureTables() error {
	s.db.MustExec(schema)
	return nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
