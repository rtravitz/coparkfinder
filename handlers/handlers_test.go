package handler

import (
	"encoding/json"
	"fmt"
	"github.com/rtravitz/coparkfinder/models"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestActivitiesIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/activities", nil)
	env := Env{DB: &mockDB{}}
	router := env.NewRouter()
	router.ServeHTTP(w, r)
	response := w.Result()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	dec := json.NewDecoder(response.Body)
	var activities []models.Activity
	err := dec.Decode(&activities)
	if err != nil {
		log.Fatal(err)
	}
	equals(t, "fishing", activities[0].Type)
	equals(t, "boating", activities[1].Type)
}

func TestFacilitiesIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/facilities", nil)
	env := Env{DB: &mockDB{}}
	router := env.NewRouter()
	router.ServeHTTP(w, r)
	response := w.Result()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	dec := json.NewDecoder(response.Body)
	var facilities []models.Facility
	err := dec.Decode(&facilities)
	if err != nil {
		log.Fatal(err)
	}
	equals(t, "visitor center", facilities[0].Type)
	equals(t, "campsites", facilities[1].Type)
}

func TestParkShowHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/park?name=Crawford", nil)
	env := Env{DB: &mockDB{}}
	router := env.NewRouter()
	router.ServeHTTP(w, r)
	response := w.Result()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var park models.Park
	err := json.NewDecoder(response.Body).Decode(&park)
	if err != nil {
		log.Fatal(err)
	}

	equals(t, park.Name, "Boyd Lake")
}

func TestAllParksHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/parks", nil)
	env := Env{DB: &mockDB{}}
	router := env.NewRouter()
	router.ServeHTTP(w, r)
	response := w.Result()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var parks []models.Park
	err := json.NewDecoder(response.Body).Decode(&parks)
	if err != nil {
		log.Fatal(err)
	}

	equals(t, parks[0].Name, "Name1")
	equals(t, parks[1].Name, "Name2")
}

type mockDB struct{}

func (mdb *mockDB) AllParks() ([]*models.Park, error) {
	facilities := make([]*models.Facility, 0)
	facilities = append(facilities, &models.Facility{1, "visitor center"})
	activities := make([]*models.Activity, 0)
	activities = append(activities, &models.Activity{1, "fishing"})
	park1 := &models.Park{
		ID:          1,
		Name:        "Name1",
		Street:      "Street1",
		City:        "City1",
		Zip:         "Zip1",
		Email:       "Email1",
		Description: "Description1",
		URL:         "Url1",
		Facilities:  facilities,
		Activities:  activities,
	}
	park2 := &models.Park{
		ID:          1,
		Name:        "Name2",
		Street:      "Street2",
		City:        "City2",
		Zip:         "Zip2",
		Email:       "Email2",
		Description: "Description2",
		URL:         "Url2",
		Facilities:  facilities,
		Activities:  activities,
	}
	parks := make([]*models.Park, 0)
	parks = append(parks, park1)
	parks = append(parks, park2)
	return parks, nil
}

func (mdb *mockDB) FindPark(where string, params ...interface{}) (*models.Park, error) {
	facilities := make([]*models.Facility, 0)
	facilities = append(facilities, &models.Facility{1, "visitor center"})
	activities := make([]*models.Activity, 0)
	activities = append(activities, &models.Activity{1, "fishing"})
	park := &models.Park{
		ID:          1,
		Name:        "Boyd Lake",
		Street:      "3720 North County Road",
		City:        "Loveland",
		Zip:         "80538",
		Email:       "boyd.lake@state.co.us",
		Description: "Colorful sailboats skimming blue water.",
		URL:         "http://cpw.state.co.us/placestogo/parks/BoydLake",
		Facilities:  facilities,
		Activities:  activities,
	}
	return park, nil
}

func (mdb *mockDB) FindParks(params map[string][]string) ([]*models.Park, error) {
	return nil, nil
}

func (mdb *mockDB) AllFacilities() ([]*models.Facility, error) {
	facilities := make([]*models.Facility, 0)
	facilities = append(facilities, &models.Facility{1, "visitor center"})
	facilities = append(facilities, &models.Facility{2, "campsites"})
	return facilities, nil
}

func (mdb *mockDB) AllActivities() ([]*models.Activity, error) {
	activities := make([]*models.Activity, 0)
	activities = append(activities, &models.Activity{1, "fishing"})
	activities = append(activities, &models.Activity{2, "boating"})
	return activities, nil
}

func (mdb *mockDB) FindParkActivities(parkID int) ([]*models.Activity, error) {
	return nil, nil
}

func (mdb *mockDB) FindParkFacilities(parkID int) ([]*models.Facility, error) {
	return nil, nil
}

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
