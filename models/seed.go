package models

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

//Seed fills the database with data imported from a CSV of parks information
func Seed(db *DB) {
	parks := unmarshalCSV()
	facilities := compileFacilitiesList(parks)
	activities := compileActivitiesList(parks)

	//insert parks
	for _, park := range parks {
		_, err := db.InsertPark(park)
		checkErr(err)
	}

	//insert facilities
	for _, facility := range facilities {
		_, err := db.InsertFacility(facility)
		checkErr(err)
	}

	//insert activities
	for _, activity := range activities {
		_, err := db.InsertActivity(activity)
		checkErr(err)
	}

	for _, park := range parks {
		dbPark, err := db.FindPark("name = $1", park.Name)
		checkErr(err)
		for _, facility := range park.facilityList {
			dbFacility, err := db.FindFacility("type = $1", facility)
			checkErr(err)
			parkFacility := ParkFacility{ParkID: dbPark.ID, FacilityID: dbFacility.ID}
			db.InsertParkFacility(parkFacility)
		}
		for _, activity := range park.activityList {
			dbActivity, err := db.FindActivity("type = $1", activity)
			checkErr(err)
			parkActivity := ParkActivity{ParkID: dbPark.ID, ActivityID: dbActivity.ID}
			db.InsertParkActivity(parkActivity)
		}
	}
}

func unmarshalCSV() []Park {
	file, err := os.Open("data/park_finder3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rawData, err := r.ReadAll()
	var park Park
	var parkList []Park

	for _, record := range rawData {
		park.Name = record[0]
		park.Street = record[1]
		park.City = record[2]
		park.Zip = record[3]
		park.Email = record[4]
		park.Description = record[5]
		park.URL = record[6]
		park.facilityList = splitAndTrimList(record[7])
		park.activityList = splitAndTrimList(record[8])
		parkList = append(parkList, park)
	}
	return parkList[1:]
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func compileFacilitiesList(parks []Park) []Facility {
	var facilityTypes []string
	var facilitiesList []Facility
	var facility Facility

	for _, park := range parks {
		for _, facilityType := range park.facilityList {
			facilityTypes = appendIfMissing(facilityTypes, facilityType)
		}
	}

	for _, facilityType := range facilityTypes {
		facility.Type = facilityType
		facilitiesList = append(facilitiesList, facility)
	}
	return facilitiesList
}

func compileActivitiesList(parks []Park) []Activity {
	var activityTypes []string
	var activitiesList []Activity
	var activity Activity

	for _, park := range parks {
		for _, activityType := range park.activityList {
			activityTypes = appendIfMissing(activityTypes, activityType)
		}
	}

	for _, activityType := range activityTypes {
		activity.Type = activityType
		activitiesList = append(activitiesList, activity)
	}
	return activitiesList
}

func appendIfMissing(current []string, toAdd string) []string {
	for _, element := range current {
		if element == toAdd {
			return current
		}
	}
	return append(current, toAdd)
}

func splitAndTrimList(list string) (cleaned []string) {
	split := strings.Split(list, ",")
	for _, element := range split {
		trimmed := strings.Trim(element, " ")
		cleaned = append(cleaned, trimmed)
	}
	return
}
