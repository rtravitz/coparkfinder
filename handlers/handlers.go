package handler

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

type Handler struct {
	DB *models.DB
}

func NewHandler(db *models.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/parks", h.ParksIndex).
		Methods("GET")
	r.HandleFunc("/api/v1/activities", h.ActivitiesIndex).
		Methods("GET")
	r.HandleFunc("/api/v1/facilities", h.FacilitiesIndex).
		Methods("GET")

	handler := handlers.HTTPMethodOverrideHandler(r)
	o := handlers.AllowedOrigins([]string{"*"})
	handler = handlers.CORS(o)(r)
	return handler
}

func (h *Handler) ActivitiesIndex(w http.ResponseWriter, r *http.Request) {
	var activities []*models.Activity
	tx, err := h.DB.Begin()
	activities, err = tx.AllActivities()
	tx.Commit()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(activities); err != nil {
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
	if err = json.NewEncoder(w).Encode(facilities); err != nil {
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
	if err = json.NewEncoder(w).Encode(parks); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
