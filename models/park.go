package models

import (
	"database/sql"
	"fmt"
	"github.com/rtravitz/coparkfinder/db"
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
	Id          int
	Name        string
	Street      string
	City        string
	Zip         string
	Email       string
	Description string
	Url         string
	Facility    string
	Activity    string
}

func InsertPark(park Park) (sql.Result, error) {
	return db.DBCon.Exec(
		fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5, $6, $7)",
			ParkTableName, ParkNameCol, ParkStreetCol, ParkCityCol, ParkZipCol,
			ParkEmailCol, ParkDescCol, ParkURLCol),
		park.Name, park.Street, park.City,
		park.Zip, park.Email, park.Description, park.Url,
	)
}
