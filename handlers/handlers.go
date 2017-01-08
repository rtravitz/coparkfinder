package handler

import (
	"encoding/json"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

func (h *Handler) ActivitiesIndex(w http.ResponseWriter, r *http.Request) {
	var activities []*models.Activity
	tx, err := h.DB.Begin()
	activities, err = tx.AllActivities()
	tx.Commit()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&activities); err != nil {
		panic(err)
	}
}

func (h *Handler) FacilitiesIndex(w http.ResponseWriter, r *http.Request) {
	var facilities []*models.Facility
	tx, err := h.DB.Begin()
	facilities, err = tx.AllFacilities()
	tx.Commit()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&facilities); err != nil {
		panic(err)
	}
}

func (h *Handler) ParksIndex(w http.ResponseWriter, r *http.Request) {
	var parks []*models.Park
	params := r.URL.Query()
	tx, err := h.DB.Begin()
	if len(params) > 0 {
		parks, err = tx.FindParks(params)
	} else {
		parks, err = tx.AllParks()
	}
	tx.Commit()
	checkErr(err)
	for _, park := range parks {
		park.Facilities, err = park.FindParkFacilities(h.DB)
		park.Activities, err = park.FindParkActivities(h.DB)
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
		tx, err := h.DB.Begin()
		park, err = tx.FindPark("name = $1", name[0])
		tx.Commit()
		checkErr(err)
	}
	park.Facilities, err = park.FindParkFacilities(h.DB)
	park.Activities, err = park.FindParkActivities(h.DB)
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
