package models

import (
	"database/sql"
	"fmt"
)

const (
	FacilityTableName = "facilities"
	FacilityTypeCol   = "type"
)

type Facility struct {
	ID   int
	Type string
}

func (tx *Tx) InsertFacility(facility Facility) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s) VALUES($1)", FacilityTableName, FacilityTypeCol),
		facility.Type,
	)
}
