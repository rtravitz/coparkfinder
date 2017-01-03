package models

import (
	"database/sql"
	"fmt"
)

//Name of the parks_activities table and associated columns in the database
const (
	ParkActivityTableName     = "parks_activities"
	ParkActivityParkIDCol     = "park_id"
	ParkActivityActivityIDCol = "activity_id"
)

//ParkActivity is a representation of the parks_activities table in the database
type ParkActivity struct {
	ID         int
	ParkID     int
	ActivityID int
}

//InsertParkActivity inserts a Facility into the database
func (tx *Tx) InsertParkActivity(parkActivity ParkActivity) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2)",
			ParkActivityTableName, ParkActivityParkIDCol, ParkActivityActivityIDCol),
		parkActivity.ParkID, parkActivity.ActivityID,
	)
}
