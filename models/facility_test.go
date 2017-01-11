package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertFacility(t *testing.T) {
	buildDB()
	defer teardownDB()
	tx, err := tdb.Begin()
	facility := newTestFacility()
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
	buildDB()
	defer teardownDB()
	_, err := insertTestFacilities()
	checkErr(err)
	tx, err := tdb.Begin()
	facility, err := tx.FindFacility("type = $1", "boathouse")
	tx.Commit()
	ok(t, err)

	equals(t, "boathouse", facility.Type)
}

func TestAllFacilities(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestFacilities()
	checkErr(err)
	tx, err := tdb.Begin()
	facilities, err := tx.AllFacilities()
	tx.Commit()
	ok(t, err)

	equals(t, "boathouse", facilities[0].Type)
	equals(t, "picnic shelter", facilities[1].Type)
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
