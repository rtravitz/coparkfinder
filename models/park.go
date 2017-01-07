package models

import (
	"database/sql"
	"fmt"
	"log"
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
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Street       string `json:"street"`
	City         string `json:"city"`
	Zip          string `json:"zip"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	URL          string `json:"url"`
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
func (tx *Tx) AllParks() ([]*Park, error) {
	rows, err := tx.Query("SELECT * FROM parks")
	if err != nil {
		return nil, err
	}
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

func (tx *Tx) FindPark(where string, params ...interface{}) (*Park, error) {
	park := new(Park)
	row := tx.QueryRow(fmt.Sprintf("SELECT * FROM parks WHERE %s", where), params...)
	err := row.Scan(&park.ID, &park.Name, &park.Street, &park.City,
		&park.Zip, &park.Email, &park.Description, &park.URL)
	if err != nil {
		return nil, err
	}
	return park, nil
}

func (tx *Tx) FindParks(params map[string][]string) ([]*Park, error) {
	rows, err := tx.Query(queryDecision(params))
	if err != nil {
		return nil, err
	}
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
	log.Print(facOK)
	log.Print(len(params) == 2)

	if len(params) == 2 {
		return allParamsQuery(facList, actList, facLen, actLen)
	}
	if facOK {
		return facilitiesQuery(facList, facLen)
	} else {
		return activitiesQuery(actList, actLen)
	}
}

func allParamsQuery(facs string, acts string, facLen int, actLen int) (query string) {
	query = fmt.Sprintf(
		`SELECT * FROM parks WHERE parks.id =
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
	log.Print(query)
	return
}
