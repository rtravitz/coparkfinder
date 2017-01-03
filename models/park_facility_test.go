package models_test

import (
	"fmt"
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertParkFacility(t *testing.T) {
	tx, err := tdb.Begin()
	parkFacility := newParkFacility()
	defer tearDown("parks_facilities", "park_id = $1 AND facility_id = $2",
		parkFacility.ParkID, parkFacility.FacilityID)
	_, err := tx.InsertParkFacility(parkFacility)
	ok(t, err)
	tx.Commit()

	row := tdb.QueryRow("SELECT * FROM parks_facilities WHERE park_id = $1 AND facility_id = $2",
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
