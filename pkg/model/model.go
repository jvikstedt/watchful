package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB(driverName string, dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func EnsureTables(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	return err
}
