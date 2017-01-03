package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertActivity(t *testing.T) {
	tx, err := tdb.Begin()
	activity := newTestActivity()
	defer tearDown("activities", "type = $1", activity.Type)
	_, err := tx.InsertActivity(activity)
	tx.Commit()
	ok(t, err)
	row := tdb.QueryRow("SELECT type FROM activities WHERE type = $1", activity.Type)
	var returnedType string
	if err := row.Scan(&returnedType); err != nil {
		panic(err)
	}

	equals(t, activity.Type, returnedType)
}

func newTestActivity() Activity {
	return Activity{Type: "fishing"}
}
