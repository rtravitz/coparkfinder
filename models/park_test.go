package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"strconv"
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
	_, err := insertTestParks(2)
	checkErr(err)
	tx, err := tdb.Begin()
	parks, err := tx.AllParks()
	tx.Commit()
	ok(t, err)

	equals(t, "Name1", parks[1].Name)
	equals(t, 2, len(parks))
}

func TestFindPark(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestParks(3)
	checkErr(err)
	tx, err := tdb.Begin()
	park, err := tx.FindPark("name = $1", "Name1")
	tx.Commit()
	ok(t, err)

	equals(t, "Name1", park.Name)
}

func TestFindParkActivities(t *testing.T) {
	buildDB()
	defer teardownDB()
	parkIds, err := insertTestParks(2)
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
	parkIds, err := insertTestParks(2)
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
	buildDB()
	defer teardownDB()
	facQuery := map[string][]string{"activities": {"'fishing'"}}
	actQuery := map[string][]string{"facilities": {"'boathouse'"}}
	bothQuery := map[string][]string{"activities": {"'fishing'"}, "facilities": {"'boathouse'"}}
	parkIds, err := insertTestParks(3)
	activityId, err := insertTestActivities()
	facilityId, err := insertTestFacilities()
	tx, err := tdb.Begin()
	_, err = tx.InsertParkFacility(ParkFacility{ParkID: parkIds[0], FacilityID: facilityId[0]})
	checkErr(err)
	_, err = tx.InsertParkFacility(ParkFacility{ParkID: parkIds[1], FacilityID: facilityId[0]})
	checkErr(err)
	_, err = tx.InsertParkActivity(ParkActivity{ParkID: parkIds[1], ActivityID: activityId[0]})
	checkErr(err)
	_, err = tx.InsertParkActivity(ParkActivity{ParkID: parkIds[2], ActivityID: activityId[0]})
	checkErr(err)
	facResult, err := tx.FindParks(facQuery)
	ok(t, err)
	actResult, err := tx.FindParks(actQuery)
	ok(t, err)
	bothResult, err := tx.FindParks(bothQuery)
	ok(t, err)
	tx.Commit()

	equals(t, 2, len(facResult))
	equals(t, 2, len(actResult))
	equals(t, 1, len(bothResult))
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

func insertTestParks(num int) (ids []int, err error) {
	var parks []Park
	for i := 0; i < num; i++ {
		var park Park
		strNum := strconv.Itoa(i)
		park = Park{
			Name:        "Name" + strNum,
			Street:      "Street" + strNum,
			City:        "City" + strNum,
			Zip:         "Zip" + strNum,
			Email:       "Email" + strNum,
			Description: "Description" + strNum,
			URL:         "URL" + strNum,
		}
		parks = append(parks, park)
	}
	tx, err := tdb.Begin()
	for _, park := range parks {
		_, err = tx.InsertPark(park)
	}
	tx.Commit()
	if err != nil {
		return nil, err
	}
	rows, err := tdb.Query("SELECT id FROM parks")
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
		ids = append(ids, id)
	}
	return ids, nil
}
