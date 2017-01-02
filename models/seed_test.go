package models

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestAppendIfMissing(t *testing.T) {
	testList := []string{"dogs", "cats", "pigs"}
	testAdd1 := "cats"
	testAdd2 := "horses"

	testList = appendIfMissing(testList, testAdd1)
	equals(t, []string{"dogs", "cats", "pigs"}, testList)

	testList = appendIfMissing(testList, testAdd2)
	equals(t, []string{"dogs", "cats", "pigs", "horses"}, testList)
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
