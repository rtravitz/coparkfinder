package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//DB is a wrapper for *sql.DB
type DB struct {
	*sql.DB
}

//Tx is a wrapper for *sql.Tx
type Tx struct {
	*sql.Tx
}

//OpenDB is a wrapper for sql.Open and returns a DB
func OpenDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

//Begin initiates a transaction
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}
