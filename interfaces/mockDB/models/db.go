package models

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

// DataStore the main place that stores the data.
type DataStore interface {
	AllBooks() ([]*Book, error)
}

// DB My main db struct
type DB struct {
	*sql.DB
}

//NewDB DB constructor
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mssql", dataSourceName)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
