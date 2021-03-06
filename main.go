package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rtravitz/coparkfinder/handlers"
	"github.com/rtravitz/coparkfinder/models"
)

func main() {
	port := os.Getenv("PORT")
	dbAddress := os.Getenv("PARKFINDER_DB")
	db, err := models.OpenDB(dbAddress)
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
	seedIfEmpty(db)

	env := &handler.Env{db}
	r := env.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func seedIfEmpty(db *models.DB) {
	parks, err := db.AllParks()
	checkErr(err)
	if len(parks) == 0 {
		models.Seed(db)
	}
}
