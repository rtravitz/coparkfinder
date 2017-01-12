package handler

import (
	"encoding/json"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

func (env *Env) ActivitiesIndex(w http.ResponseWriter, r *http.Request) {
	var activities []*models.Activity
	activities, err := env.DB.AllActivities()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&activities); err != nil {
		panic(err)
	}
}

func (env *Env) FacilitiesIndex(w http.ResponseWriter, r *http.Request) {
	var facilities []*models.Facility
	facilities, err := env.DB.AllFacilities()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&facilities); err != nil {
		panic(err)
	}
}

func (env *Env) ParksIndex(w http.ResponseWriter, r *http.Request) {
	var parks []*models.Park
	var err error
	params := r.URL.Query()
	if len(params) > 0 {
		parks, err = env.DB.FindParks(params)
	} else {
		parks, err = env.DB.AllParks()
	}
	checkErr(err)
	for _, park := range parks {
		park.Facilities, err = env.DB.FindParkFacilities(park.ID)
		park.Activities, err = env.DB.FindParkActivities(park.ID)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&parks); err != nil {
		panic(err)
	}
}

func (env *Env) ParkShow(w http.ResponseWriter, r *http.Request) {
	var park *models.Park
	var err error
	params := r.URL.Query()
	if name, ok := params["name"]; ok {
		park, err = env.DB.FindPark("name = $1", name[0])
		checkErr(err)
	}
	park.Facilities, err = env.DB.FindParkFacilities(park.ID)
	park.Activities, err = env.DB.FindParkActivities(park.ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&park); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
