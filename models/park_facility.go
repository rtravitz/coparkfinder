package models

import (
	"database/sql"
	"fmt"
)

//Name of the parks_facilities table and associated columns in the database
const (
	ParkFacilityTableName     = "parks_facilities"
	ParkFacilityParkIDCol     = "park_id"
	ParkFacilityFacilityIDCol = "facility_id"
)

//ParkFacility is a representation of the parks_facilities table in the database
type ParkFacility struct {
	ID         int
	ParkID     int
	FacilityID int
}

//InsertParkFacility inserts a Facility into the database
func (db *DB) InsertParkFacility(parkFacility ParkFacility) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2)",
			ParkFacilityTableName, ParkFacilityParkIDCol, ParkFacilityFacilityIDCol),
		parkFacility.ParkID, parkFacility.FacilityID,
	)
}
