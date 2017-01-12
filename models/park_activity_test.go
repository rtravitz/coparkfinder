package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertParkActivity(t *testing.T) {
	buildDB()
	defer teardownDB()
	parkActivity := newParkActivity()
	_, err := tdb.InsertParkActivity(parkActivity)
	ok(t, err)

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
