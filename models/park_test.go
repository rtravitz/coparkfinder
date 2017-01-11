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

func TestFindParkActivities(t *testing.T) {
	parkIds, err := insertTestParks()
	activityIds, err := insertTestActivities()
	testParkActivity := ParkActivity{ParkID: parkIds[0], ActivityID: activityIds[0]}
	defer tearDown("parks", "id IN ($1, $2)", parkIds[0], parkIds[1])
	defer tearDown("activities", "id IN ($1, $2)", activityIds[0], activityIds[1])
	defer tearDown("parks_activities", "park_id = $1", parkIds[0])
	tx, err := tdb.Begin()
	_, err = tx.InsertParkActivity(testParkActivity)
	tx.Commit()
	testPark := Park{ID: parkIds[0]}
	activities, err := testPark.FindParkActivities(tdb)
	ok(t, err)

	equals(t, 1, len(activities))
	equals(t, "fishing", activities[0].Type)
}

func TestFindParkFacilities(t *testing.T) {
	parkIds, err := insertTestParks()
	facilityIds, err := insertTestFacilities()
	testParkFacility := ParkFacility{ParkID: parkIds[0], FacilityID: facilityIds[0]}
	defer tearDown("parks", "id IN ($1, $2)", parkIds[0], parkIds[1])
	defer tearDown("facilities", "id IN ($1, $2)", facilityIds[0], facilityIds[1])
	defer tearDown("parks_facilities", "park_id = $1", parkIds[0])
	tx, err := tdb.Begin()
	_, err = tx.InsertParkFacility(testParkFacility)
	tx.Commit()
	testPark := Park{ID: parkIds[0]}
	facilities, err := testPark.FindParkFacilities(tdb)
	ok(t, err)

	equals(t, 1, len(facilities))
	equals(t, "boathouse", facilities[0].Type)
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
