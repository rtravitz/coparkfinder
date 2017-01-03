package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertActivity(t *testing.T) {
	tx, err := tdb.Begin()
	activity := newTestActivity()
	defer tearDown("activities", "type = $1", activity.Type)
	_, err = tx.InsertActivity(activity)
	tx.Commit()
	ok(t, err)
	row := tdb.QueryRow("SELECT type FROM activities WHERE type = $1", activity.Type)
	var returnedType string
	if err := row.Scan(&returnedType); err != nil {
		panic(err)
	}

	equals(t, activity.Type, returnedType)
}

func TestFindActivity(t *testing.T) {
	ids, err := insertTestActivities()
	checkErr(err)
	defer tearDown("activities", "id IN ($1, $2)", ids[0], ids[1])
	tx, err := tdb.Begin()
	activity, err := tx.FindActivity("type = $1", "fishing")
	tx.Commit()
	ok(t, err)

	equals(t, "fishing", activity.Type)
}

func newTestActivity() Activity {
	return Activity{Type: "fishing"}
}

func insertTestActivities() (teardownIDs []int, err error) {
	tx, err := tdb.Begin()
	activity1 := newTestActivity()
	activity2 := Activity{Type: "hiking"}
	activitiesList := []Activity{activity1, activity2}
	for _, activity := range activitiesList {
		_, err = tx.InsertActivity(activity)
	}
	tx.Commit()
	rows, err := tdb.Query("SELECT id FROM activities")
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
