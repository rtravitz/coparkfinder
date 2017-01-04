package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rtravitz/coparkfinder/models"
)

func main() {
	db, err := models.OpenDB(dbinfo())
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
	seedIfEmpty(db)

	fmt.Println("Successfully connected!")
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

func dbinfo() string {
	port, err := strconv.Atoi(os.Getenv("PARKFINDER_PORT"))
	checkErr(err)
	p := psqlInfo{host: os.Getenv("PARKFINDER_HOST"), port: port,
		user: os.Getenv("PARKFINDER_USER"), dbname: os.Getenv("PARKFINDER_DBNAME")}
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.dbname)
}

type psqlInfo struct {
	host   string
	port   int
	user   string
	dbname string
}
