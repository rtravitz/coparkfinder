package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rtravitz/coparkfinder/handlers"
	"github.com/rtravitz/coparkfinder/models"
)

func main() {
	dbAddress := os.Getenv("PARKFINDER_DB")
	db, err := models.OpenDB(dbAddress)
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
	seedIfEmpty(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/parks", handlers.ParksIndex(db)).
		Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func seedIfEmpty(db *models.DB) {
	tx, err := db.Begin()
	parks, err := tx.AllParks()
	checkErr(err)
	tx.Commit()
	if len(parks) == 0 {
		models.Seed(db)
	}
}
