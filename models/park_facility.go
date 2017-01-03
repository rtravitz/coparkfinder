package models

import (
	"database/sql"
	"fmt"
)

const (
	ParkFacilityTableName     = "parks_facilities"
	ParkFacilityParkIDCol     = "park_id"
	ParkFacilityFacilityIDCol = "facility_id"
)

type ParkFacility struct {
	ID         int
	ParkID     int
	FacilityID int
}

func (tx *Tx) InsertParkFacility(parkFacility ParkFacility) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2)",
			ParkFacilityTableName, ParkFacilityParkIDCol, ParkFacilityFacilityIDCol),
		parkFacility.ParkID, parkFacility.FacilityID,
	)
}
