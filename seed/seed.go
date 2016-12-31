package seed

import (
	"encoding/csv"
	"log"
	"os"
)

type Park struct {
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

func UnmarshalCSV() []Park {
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
