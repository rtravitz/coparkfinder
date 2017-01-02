package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func Seed(db *DB) {
	parks := unmarshalCSV()
	facilities := compileFacilitiesList(parks)
	tx, _ := db.Begin()
	for _, park := range parks {
		result, err := tx.InsertPark(park)
		checkErr(err)
		fmt.Println(result)
	}
	tx.Commit()

	tx, _ = db.Begin()
	for _, facility := range facilities {
		result, err := tx.InsertFacility(facility)
		checkErr(err)
		fmt.Println(result)
	}
	tx.Commit()
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
		park.Url = record[6]
		park.facilityList = splitAndTrimList(record[7])
		park.activityList = record[8]
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
	fmt.Println(facilityTypes)

	for _, facilityType := range facilityTypes {
		facility.Type = facilityType
		facilitiesList = append(facilitiesList, facility)
	}
	return facilitiesList
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
