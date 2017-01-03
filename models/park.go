package models

import (
	"database/sql"
	"fmt"
)

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

type Park struct {
	ID           int
	Name         string
	Street       string
	City         string
	Zip          string
	Email        string
	Description  string
	Url          string
	facilityList []string
	activityList string
}

func (tx *Tx) InsertPark(park Park) (sql.Result, error) {
	return tx.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5, $6, $7)",
			ParkTableName, ParkNameCol, ParkStreetCol,
			ParkCityCol, ParkZipCol, ParkEmailCol,
			ParkDescCol, ParkURLCol),
		park.Name, park.Street, park.City,
		park.Zip, park.Email, park.Description, park.Url,
	)
}

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
			&park.Zip, &park.Email, &park.Description, &park.Url)
		if err != nil {
			return nil, err
		}
		parks = append(parks, park)
	}
	return parks, nil
}
