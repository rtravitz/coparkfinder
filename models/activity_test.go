package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertActivity(t *testing.T) {
	buildDB()
	defer teardownDB()
	activity := newTestActivity()
	_, err := tdb.InsertActivity(activity)
	ok(t, err)
	row := tdb.QueryRow("SELECT type FROM activities WHERE type = $1", activity.Type)
	var returnedType string
	if err := row.Scan(&returnedType); err != nil {
		panic(err)
	}

	equals(t, activity.Type, returnedType)
}

func TestFindActivity(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestActivities()
	checkErr(err)
	activity, err := tdb.FindActivity("type = $1", "fishing")
	ok(t, err)

	equals(t, "fishing", activity.Type)
}

func TestAllActivities(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestActivities()
	checkErr(err)
	activities, err := tdb.AllActivities()
	ok(t, err)

	equals(t, "fishing", activities[0].Type)
	equals(t, "hiking", activities[1].Type)
}

func newTestActivity() Activity {
	return Activity{Type: "fishing"}
}

func insertTestActivities() (teardownIDs []int, err error) {
	activity1 := newTestActivity()
	activity2 := Activity{Type: "hiking"}
	activitiesList := []Activity{activity1, activity2}
	for _, activity := range activitiesList {
		_, err = tdb.InsertActivity(activity)
	}
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
