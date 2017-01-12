package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//DB is a wrapper for *sql.DB
type DB struct {
	*sql.DB
}

type Datastore interface {
	AllParks() ([]*Park, error)
	FindPark(where string, params ...interface{}) (*Park, error)
	FindParks(params map[string][]string) ([]*Park, error)
	FindParkActivities(parkID int) ([]*Activity, error)
	FindParkFacilities(parkID int) ([]*Facility, error)
	AllFacilities() ([]*Facility, error)
	AllActivities() ([]*Activity, error)
}

//OpenDB is a wrapper for sql.Open and returns a DB
func OpenDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
