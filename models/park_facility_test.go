package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertParkFacility(t *testing.T) {
	buildDB()
	defer teardownDB()
	parkFacility := newParkFacility()
	_, err := tdb.InsertParkFacility(parkFacility)
	ok(t, err)

	row := tdb.QueryRow("SELECT park_id, facility_id FROM parks_facilities WHERE park_id = $1 AND facility_id = $2",
		parkFacility.ParkID, parkFacility.FacilityID)
	var returnedParkID, returnedFacilityID int
	if err := row.Scan(&returnedParkID, &returnedFacilityID); err != nil {
		panic(err)
	}

	equals(t, parkFacility.ParkID, returnedParkID)
	equals(t, parkFacility.FacilityID, returnedFacilityID)
}

func newParkFacility() ParkFacility {
	return ParkFacility{ParkID: 1, FacilityID: 1}
}
