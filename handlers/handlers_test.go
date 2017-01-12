package handler

import (
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivitiesIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/activities", nil)
	env := Env{DB: &mockDB{}}
	router := env.NewRouter()
	router.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

type mockDB struct{}

func (mdb *mockDB) AllParks() ([]*models.Park, error) {
	return nil, nil
}

func (mdb *mockDB) FindPark(where string, params ...interface{}) (*models.Park, error) {
	return nil, nil
}

func (mdb *mockDB) FindParks(params map[string][]string) ([]*models.Park, error) {
	return nil, nil
}

func (mdb *mockDB) FindParkActivities(parkID int) ([]*models.Activity, error) {
	return nil, nil
}

func (mdb *mockDB) FindParkFacilities(parkID int) ([]*models.Facility, error) {
	return nil, nil
}

func (mdb *mockDB) AllFacilities() ([]*models.Facility, error) {
	return nil, nil
}

func (mdb *mockDB) AllActivities() ([]*models.Activity, error) {
	return nil, nil
}
