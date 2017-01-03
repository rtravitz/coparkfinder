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

func TestFindFacility(t *testing.T) {
	ids, err := insertTestFacilities()
	checkErr(err)
	defer tearDown("facilities", "id IN ($1, $2)", ids[0], ids[1])
	tx, err := tdb.Begin()
	facility, err := tx.FindFacility("type = $1", "boathouse")

	equals(t, "boathouse", facility.Type)
}

func newTestFacility() Facility {
	return Facility{Type: "boathouse"}
}

func insertTestFacilities() (teardownIDs []int, err error) {
	tx, err := tdb.Begin()
	facility1 := newTestFacility()
	facility2 := Facility{Type: "picnic shelter"}
	facilitiesList := []Facility{facility1, facility2}
	for _, facility := range facilitiesList {
		_, err = tx.InsertFacility(facility)
	}
	tx.Commit()
	rows, err := tdb.Query("SELECT id FROM facilities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		teardownIDs = append(teardownIDs, id)
	}
	return teardownIDs, nil
}
