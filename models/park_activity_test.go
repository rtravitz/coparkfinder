package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertParkActivity(t *testing.T) {
	tx, err := tdb.Begin()
	parkActivity := newParkActivity()
	defer tearDown("parks_activities", "park_id = $1 AND activity_id = $2",
		parkActivity.ParkID, parkActivity.ActivityID)
	_, err = tx.InsertParkActivity(parkActivity)
	ok(t, err)
	tx.Commit()

	row := tdb.QueryRow("SELECT park_id, activity_id FROM parks_activities WHERE park_id = $1 AND activity_id = $2",
		parkActivity.ParkID, parkActivity.ActivityID)
	var returnedParkID, returnedActivityID int
	if err := row.Scan(&returnedParkID, &returnedActivityID); err != nil {
		panic(err)
	}

	equals(t, parkActivity.ParkID, returnedParkID)
	equals(t, parkActivity.ActivityID, returnedActivityID)
}

func newParkActivity() ParkActivity {
	return ParkActivity{ParkID: 1, ActivityID: 1}
}
