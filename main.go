package main

import (
	"fmt"

	"github.com/rtravitz/coparkfinder/db"
	"github.com/rtravitz/coparkfinder/seed"
)

func main() {
	var err error
	db.DBCon, err = db.OpenDB()
	checkErr(err)
	defer db.DBCon.Close()
	err = db.DBCon.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	seed.Seed()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
