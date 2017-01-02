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

func TestInsertPark(t *testing.T) {
	tx, err := tdb.Begin()
	park := newTestPark()
	defer tearDown("parks", "name = $1", park.Name)
	_, err = tx.InsertPark(park)
	ok(t, err)
	tx.Commit()

	row := tdb.QueryRow("SELECT city FROM parks WHERE name = $1", park.Name)
	var returnedCity string
	if err := row.Scan(&returnedCity); err != nil {
		panic(err)
	}

	equals(t, "Loveland", returnedCity)
}

func newTestPark() Park {
	return Park{
		Name:        "Boyd Lake",
		Street:      "3720 North County Road",
		City:        "Loveland",
		Zip:         "80538",
		Email:       "boyd.lake@state.co.us",
		Description: "Colorful sailboats skimming blue water.",
		Url:         "http://cpw.state.co.us/placestogo/parks/BoydLake",
	}
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
