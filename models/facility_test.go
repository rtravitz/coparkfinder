package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertFacility(t *testing.T) {
	tx, err := tdb.Begin()
	facility := newTestFacility()
	defer tearDown("facilities", "type = $1", facility.Type)
	_, err = tx.InsertFacility(facility)
	ok(t, err)
	tx.Commit()

	row := tdb.QueryRow("SELECT type FROM facilities WHERE type = $1", facility.Type)
	var returnedType string
	if err := row.Scan(&returnedType); err != nil {
		panic(err)
	}

	equals(t, facility.Type, returnedType)
}

func newTestFacility() Facility {
	return Facility{Type: "boathouse"}
}
