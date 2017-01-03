package models

import (
	"database/sql"
	"fmt"
)

//Name of the facilities table and associated columns in the database
const (
	FacilityTableName = "facilities"
	FacilityTypeCol   = "type"
)

//Facility is a representation of the facilities table in the database
type Facility struct {
	ID   int
	Type string
}

//InsertFacility inserts a Facility into the database
func (tx *Tx) InsertFacility(facility Facility) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s) VALUES($1)", FacilityTableName, FacilityTypeCol),
		facility.Type,
	)
}
