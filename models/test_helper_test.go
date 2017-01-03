package models_test

import (
	"fmt"
	. "github.com/rtravitz/coparkfinder/models"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var tdb *DB

func init() {
	var err error
	source := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		"localhost", 5432, "rtravitz", "parkfindertest")
	tdb, err = OpenDB(source)
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertTestParks() (teardownIDs []int, err error) {
	tx, err := tdb.Begin()
	park1 := newTestPark()
	park2 := newTestPark()
	parkList := []Park{park1, park2}
	for _, park := range parkList {
		_, err = tx.InsertPark(park)
	}
	tx.Commit()
	rows, err := tdb.Query("SELECT id FROM parks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		teardownIDs = append(teardownIDs, id)
	}
	return teardownIDs, nil
}

func tearDown(table, where string, params ...interface{}) {
	_, err := tdb.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE %s", table, where),
		params...)
	if err != nil {
		panic(fmt.Sprintf("Problem tearing down %s data: %v", table, err))
	}
}

// assert fails the test if the condition is false
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
