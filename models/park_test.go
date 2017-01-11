package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestInsertPark(t *testing.T) {
	buildDB()
	defer teardownDB()
	tx, err := tdb.Begin()
	park := newTestPark()
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
	buildDB()
	defer teardownDB()
	_, err := insertTestParks()
	checkErr(err)
	tx, err := tdb.Begin()
	parks, err := tx.AllParks()
	tx.Commit()
	ok(t, err)

	equals(t, "Boyd Lake", parks[1].Name)
	equals(t, 2, len(parks))
}

func TestFindPark(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestParks()
	checkErr(err)
	tx, err := tdb.Begin()
	park, err := tx.FindPark("name = $1", "Boyd Lake")
	tx.Commit()
	ok(t, err)

	equals(t, "Boyd Lake", park.Name)
}

func TestFindParkActivities(t *testing.T) {
	buildDB()
	defer teardownDB()
	parkIds, err := insertTestParks()
	activityIds, err := insertTestActivities()
	testParkActivity := ParkActivity{ParkID: parkIds[0], ActivityID: activityIds[0]}
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
	buildDB()
	defer teardownDB()
	parkIds, err := insertTestParks()
	facilityIds, err := insertTestFacilities()
	testParkFacility := ParkFacility{ParkID: parkIds[0], FacilityID: facilityIds[0]}
	tx, err := tdb.Begin()
	_, err = tx.InsertParkFacility(testParkFacility)
	tx.Commit()
	testPark := Park{ID: parkIds[0]}
	facilities, err := testPark.FindParkFacilities(tdb)
	ok(t, err)

	equals(t, 1, len(facilities))
	equals(t, "boathouse", facilities[0].Type)
}

func TestFindParks(t *testing.T) {

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
