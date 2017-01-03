package models

import (
	"database/sql"
	"fmt"
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
	ID           int
	Name         string
	Street       string
	City         string
	Zip          string
	Email        string
	Description  string
	URL          string
	facilityList []string
	activityList string
}

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
