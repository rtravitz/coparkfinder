package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertPark(t *testing.T) {
	tx, err := tdb.Begin()
	park := newTestPark()
	defer tearDown("parks", "name = $1", park.Name)
	_, err = tx.InsertPark(park)
	ok(t, err)
	tx.Commit()

	row := tdb.QueryRow("SELECT city FROM parks WHERE name = $1", park.Name)
	var returnedCity string
	if err := row.Scan(&returnedCity); err != nil {
		panic(err)
	}

	equals(t, park.City, returnedCity)
}

func TestAllParks(t *testing.T) {
	ids, err := insertTestParks()
	checkErr(err)
	defer tearDown("parks", "id IN ($1, $2)", ids[0], ids[1])
	tx, err := tdb.Begin()
	parks, err := tx.AllParks()
	tx.Commit()
	ok(t, err)

	equals(t, "Boyd Lake", parks[1].Name)
	equals(t, 2, len(parks))
}

func TestFindPark(t *testing.T) {
	ids, err := insertTestParks()
	checkErr(err)
	defer tearDown("parks", "id IN ($1, $2)", ids[0], ids[1])
	tx, err := tdb.Begin()
	park, err := tx.FindPark("name = $1", "Boyd Lake")
	tx.Commit()
	ok(t, err)

	equals(t, "Boyd Lake", park.Name)
}

func newTestPark() Park {
	return Park{
		Name:        "Boyd Lake",
		Street:      "3720 North County Road",
		City:        "Loveland",
		Zip:         "80538",
		Email:       "boyd.lake@state.co.us",
		Description: "Colorful sailboats skimming blue water.",
		URL:         "http://cpw.state.co.us/placestogo/parks/BoydLake",
	}
}
