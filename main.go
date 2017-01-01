package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rtravitz/coparkfinder/db"
	"github.com/rtravitz/coparkfinder/seed"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PARKFINDER_PORT"))
	checkErr(err)
	psql := psqlInfo{host: os.Getenv("PARKFINDER_HOST"), port: port,
		user: os.Getenv("PARKFINDER_USER"), dbname: os.Getenv("PARKFINDER_DBNAME")}
	db.DBCon, err = sql.Open("postgres", psql.info())
	checkErr(err)
	defer db.DBCon.Close()
	err = db.DBCon.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	seed.Seed()
}

type psqlInfo struct {
	host   string
	port   int
	user   string
	dbname string
}

func (p psqlInfo) info() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.dbname)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
