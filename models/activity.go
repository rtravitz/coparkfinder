package models

import (
	"database/sql"
	"fmt"
)

//Name of the activities table and associated columns in the database
const (
	ActivityTableName = "activities"
	ActivityTypeCol   = "type"
)

//Activity is a representation of the activities table in the database
type Activity struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

//InsertActivity inserts a Activity into the database
func (db *DB) InsertActivity(activity Activity) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s(%s) VALUES($1)", ActivityTableName, ActivityTypeCol),
		activity.Type,
	)
}

//AllActivities finds and returns all Activities from the database
func (db *DB) AllActivities() ([]*Activity, error) {
	rows, err := db.Query("SELECT * FROM activities ORDER BY type")
	if err != nil {
		return nil, err
	}
	return generateActivities(rows)
}

//FindActivity finds and returns a single Activity matching the params from the database
func (db *DB) FindActivity(where string, params ...interface{}) (*Activity, error) {
	activity := new(Activity)
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM activities WHERE %s", where), params...)
	err := row.Scan(&activity.ID, &activity.Type)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func generateActivities(rows *sql.Rows) ([]*Activity, error) {
	defer rows.Close()
	activities := make([]*Activity, 0)
	for rows.Next() {
		activity := new(Activity)
		err := rows.Scan(&activity.ID, &activity.Type)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}
