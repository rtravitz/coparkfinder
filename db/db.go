package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var (
	DBCon *sql.DB
)

type psqlInfo struct {
	host   string
	port   int
	user   string
	dbname string
}

func (p psqlInfo) Info() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.dbname)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenDB() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("PARKFINDER_PORT"))
	checkErr(err)
	psql := psqlInfo{host: os.Getenv("PARKFINDER_HOST"), port: port,
		user: os.Getenv("PARKFINDER_USER"), dbname: os.Getenv("PARKFINDER_DBNAME")}
	return sql.Open("postgres", psql.Info())
}
