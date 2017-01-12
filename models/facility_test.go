package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertFacility(t *testing.T) {
	buildDB()
	defer teardownDB()
	facility := newTestFacility()
	_, err := tdb.InsertFacility(facility)
	ok(t, err)

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
	facility, err := tdb.FindFacility("type = $1", "boathouse")
	ok(t, err)

	equals(t, "boathouse", facility.Type)
}

func TestAllFacilities(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestFacilities()
	checkErr(err)
	facilities, err := tdb.AllFacilities()
	ok(t, err)

	equals(t, "boathouse", facilities[0].Type)
	equals(t, "picnic shelter", facilities[1].Type)
}

func newTestFacility() Facility {
	return Facility{Type: "boathouse"}
}

func insertTestFacilities() (teardownIDs []int, err error) {
	facility1 := newTestFacility()
	facility2 := Facility{Type: "picnic shelter"}
	facilitiesList := []Facility{facility1, facility2}
	for _, facility := range facilitiesList {
		_, err = tdb.InsertFacility(facility)
	}
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
