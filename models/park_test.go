package models_test

import (
	. "github.com/rtravitz/coparkfinder/models"
	"strconv"
	"testing"
)

func TestInsertPark(t *testing.T) {
	buildDB()
	defer teardownDB()
	park := newTestPark()
	_, err := tdb.InsertPark(park)
	ok(t, err)

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
	parks, err := tdb.AllParks()
	ok(t, err)

	equals(t, "Name1", parks[1].Name)
	equals(t, 2, len(parks))
}

func TestFindPark(t *testing.T) {
	buildDB()
	defer teardownDB()
	_, err := insertTestParks(3)
	checkErr(err)
	park, err := tdb.FindPark("name = $1", "Name1")
	ok(t, err)

	equals(t, "Name1", park.Name)
}

func TestFindParkActivities(t *testing.T) {
	buildDB()
	defer teardownDB()
	parkIds, err := insertTestParks(2)
	activityIds, err := insertTestActivities()
	testParkActivity := ParkActivity{ParkID: parkIds[0], ActivityID: activityIds[0]}
	_, err = tdb.InsertParkActivity(testParkActivity)
	testPark := Park{ID: parkIds[0]}
	activities, err := tdb.FindParkActivities(testPark.ID)
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
	_, err = tdb.InsertParkFacility(testParkFacility)
	testPark := Park{ID: parkIds[0]}
	facilities, err := tdb.FindParkFacilities(testPark.ID)
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
	_, err = tdb.InsertParkFacility(ParkFacility{ParkID: parkIds[0], FacilityID: facilityId[0]})
	checkErr(err)
	_, err = tdb.InsertParkFacility(ParkFacility{ParkID: parkIds[1], FacilityID: facilityId[0]})
	checkErr(err)
	_, err = tdb.InsertParkActivity(ParkActivity{ParkID: parkIds[1], ActivityID: activityId[0]})
	checkErr(err)
	_, err = tdb.InsertParkActivity(ParkActivity{ParkID: parkIds[2], ActivityID: activityId[0]})
	checkErr(err)
	facResult, err := tdb.FindParks(facQuery)
	ok(t, err)
	actResult, err := tdb.FindParks(actQuery)
	ok(t, err)
	bothResult, err := tdb.FindParks(bothQuery)
	ok(t, err)

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
	for _, park := range parks {
		_, err = tdb.InsertPark(park)
	}
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
