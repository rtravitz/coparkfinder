package models_test

import (
	"fmt"
	. "github.com/rtravitz/coparkfinder/models"
	"testing"
)

func TestOpenDB(t *testing.T) {
	source := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		"localhost", 5432, "rtravitz", "parkfindertest")
	db, err := OpenDB(source)
	defer db.Close()
	ok(t, err)
}
