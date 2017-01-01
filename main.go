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

	fmt.Println("Successfully connected!")

	models.Seed(db)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
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
