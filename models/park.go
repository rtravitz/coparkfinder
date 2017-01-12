package models

import (
	"database/sql"
	"fmt"
	"strings"
)

//Name of the parks table and associated columns in the database
const (
	ParkTableName = "parks"
	ParkNameCol   = "name"
	ParkStreetCol = "street"
	ParkCityCol   = "city"
	ParkZipCol    = "zip"
	ParkEmailCol  = "email"
	ParkDescCol   = "description"
	ParkURLCol    = "url"
)

//Park is a representation of the parks table in the database
type Park struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Street       string      `json:"street"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Email        string      `json:"email"`
	Description  string      `json:"description"`
	URL          string      `json:"url"`
	Facilities   []*Facility `json:"facilities"`
	Activities   []*Activity `json:"activities"`
	facilityList []string
	activityList []string
}

type Parks []Park

//InsertPark inserts a Park into the database
func (tx *Tx) InsertPark(park Park) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5, $6, $7)",
			ParkTableName, ParkNameCol, ParkStreetCol,
			ParkCityCol, ParkZipCol, ParkEmailCol,
			ParkDescCol, ParkURLCol),
		park.Name, park.Street, park.City,
		park.Zip, park.Email, park.Description, park.URL,
	)
}

//AllParks queries the database and returns a slice of Parks
func (db DB) AllParks() ([]*Park, error) {
	rows, err := db.Query("SELECT * FROM parks")
	if err != nil {
		return nil, err
	}
	return generateParks(rows)
}

//FindPark returns a single park matching the query params
func (db DB) FindPark(where string, params ...interface{}) (*Park, error) {
	park := new(Park)
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM parks WHERE %s", where), params...)
	err := row.Scan(&park.ID, &park.Name, &park.Street, &park.City,
		&park.Zip, &park.Email, &park.Description, &park.URL)
	if err != nil {
		return nil, err
	}
	return park, nil
}

//FindParks returns a slice of parks matching the passed in query params
func (db DB) FindParks(params map[string][]string) ([]*Park, error) {
	rows, err := db.Query(queryDecision(params))
	if err != nil {
		return nil, err
	}
	return generateParks(rows)
}

//FindParkActivities returns a slice of all activities associated with a park
func (db DB) FindParkActivities(parkID int) ([]*Activity, error) {
	query := fmt.Sprintf(`SELECT activities.* FROM activities
		JOIN parks_activities ON activities.id = parks_activities.activity_id
		JOIN parks ON parks_activities.park_id = parks.id
		WHERE parks.id = %d`, parkID)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return generateActivities(rows)
}

//Finds all facilities based on a Park ID
func (db DB) FindParkFacilities(parkID int) ([]*Facility, error) {
	query := fmt.Sprintf(`SELECT facilities.* FROM facilities
		JOIN parks_facilities ON facilities.id = parks_facilities.facility_id
		JOIN parks ON parks_facilities.park_id = parks.id
		WHERE parks.id = %d`, parkID)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return generateFacilities(rows)
}

func queryDecision(params map[string][]string) string {
	var facList, actList string
	var facLen, actLen int
	var facOK bool

	if facilities, ok := params["facilities"]; ok {
		facOK = ok
		facList = facilities[0]
		facLen = len(strings.Split(facList, ","))
	}
	if activities, ok := params["activities"]; ok {
		actList = activities[0]
		actLen = len(strings.Split(actList, ","))
	}

	if len(params) == 2 {
		return allParamsQuery(facList, actList, facLen, actLen)
	}
	if facOK {
		return facilitiesQuery(facList, facLen)
	} else {
		return activitiesQuery(actList, actLen)
	}
}

func generateParks(rows *sql.Rows) ([]*Park, error) {
	defer rows.Close()
	parks := make([]*Park, 0)
	for rows.Next() {
		park := new(Park)
		err := rows.Scan(&park.ID, &park.Name, &park.Street, &park.City,
			&park.Zip, &park.Email, &park.Description, &park.URL)
		if err != nil {
			return nil, err
		}
		parks = append(parks, park)
	}
	return parks, nil
}

func allParamsQuery(facs string, acts string, facLen int, actLen int) (query string) {
	query = fmt.Sprintf(
		`SELECT * FROM parks WHERE parks.id IN
			 (SELECT fac.id FROM
					(SELECT parks.id AS id FROM parks
						JOIN parks_facilities ON parks.id = parks_facilities.park_id
						JOIN facilities ON parks_facilities.facility_id = facilities.id
						WHERE facilities.type IN (%s)
						GROUP BY parks.id
						HAVING COUNT(*) = %d) AS fac,
					(SELECT parks.id AS id FROM parks
						JOIN parks_activities ON parks.id = parks_activities.park_id
						JOIN activities ON parks_activities.activity_id = activities.id
						WHERE activities.type IN (%s)
						GROUP BY parks.id
						HAVING COUNT(*) = %d) as act
				WHERE fac.id = act.id);`,
		facs, facLen, acts, actLen)
	return
}

func activitiesQuery(acts string, actLen int) (query string) {
	query = fmt.Sprintf(
		`SELECT parks.* from parks
		JOIN parks_activities ON parks.id = parks_activities.park_id
		JOIN activities ON parks_activities.activity_id = activities.id
		WHERE activities.type IN (%s)
		GROUP BY parks.id
		HAVING COUNT(*) = %d;`,
		acts, actLen)
	return
}

func facilitiesQuery(facs string, facLen int) (query string) {
	query = fmt.Sprintf(
		`SELECT parks.* from parks
		JOIN parks_facilities ON parks.id = parks_facilities.park_id
		JOIN facilities ON parks_facilities.facility_id = facilities.id
		WHERE facilities.type IN (%s)
		GROUP BY parks.id
		HAVING COUNT(*) = %d;`,
		facs, facLen)
	return
}
