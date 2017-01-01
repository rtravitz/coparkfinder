package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rtravitz/coparkfinder/db"
	"github.com/rtravitz/coparkfinder/seed"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "rtravitz"
	dbname = "parkfinderdev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	var err error
	db.DBCon, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.DBCon.Close()
	err = db.DBCon.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	seed.Seed()
}
