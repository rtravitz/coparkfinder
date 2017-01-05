package handlers

import (
	"encoding/json"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

func ParksIndex(db *models.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin()
		parks, err := tx.AllParks()
		tx.Commit()
		checkErr(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(parks)
	}
	return http.HandlerFunc(fn)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
