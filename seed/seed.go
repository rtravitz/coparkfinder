package seed

import (
	"encoding/csv"
	"fmt"
	"github.com/rtravitz/coparkfinder/models"
	"log"
	"os"
)

func Seed() {
	parks := unmarshalCSV()
	for _, park := range parks {
		result, err := models.InsertPark(park)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}

func unmarshalCSV() []models.Park {
	file, err := os.Open("data/park_finder3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rawData, err := r.ReadAll()
	var park models.Park
	var parkList []models.Park

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
