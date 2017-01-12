package handler

import (
	"encoding/json"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

func (h *Handler) ActivitiesIndex(w http.ResponseWriter, r *http.Request) {
	var activities []*models.Activity
	activities, err := h.DB.AllActivities()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&activities); err != nil {
		panic(err)
	}
}

func (h *Handler) FacilitiesIndex(w http.ResponseWriter, r *http.Request) {
	var facilities []*models.Facility
	facilities, err := h.DB.AllFacilities()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&facilities); err != nil {
		panic(err)
	}
}

func (h *Handler) ParksIndex(w http.ResponseWriter, r *http.Request) {
	var parks []*models.Park
	var err error
	params := r.URL.Query()
	if len(params) > 0 {
		parks, err = h.DB.FindParks(params)
	} else {
		parks, err = h.DB.AllParks()
	}
	checkErr(err)
	for _, park := range parks {
		park.Facilities, err = h.DB.FindParkFacilities(park.ID)
		park.Activities, err = h.DB.FindParkActivities(park.ID)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&parks); err != nil {
		panic(err)
	}
}

func (h *Handler) ParkShow(w http.ResponseWriter, r *http.Request) {
	var park *models.Park
	var err error
	params := r.URL.Query()
	if name, ok := params["name"]; ok {
		park, err = h.DB.FindPark("name = $1", name[0])
		checkErr(err)
	}
	park.Facilities, err = h.DB.FindParkFacilities(park.ID)
	park.Activities, err = h.DB.FindParkActivities(park.ID)
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
