package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Seed(db *DB) {
	parks := unmarshalCSV()
	tx, _ := db.Begin()
	for _, park := range parks {
		result, err := tx.InsertPark(park)
		if err != nil {
			log.Fatal(err)
		}
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
		park.Facility = record[7]
		park.Activity = record[8]
		parkList = append(parkList, park)
	}
	return parkList[1:]
}
