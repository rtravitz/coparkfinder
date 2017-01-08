package handler

import (
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
	r.HandleFunc("/api/v1/park", h.ParkShow).
		Methods("GET")

	handler := handlers.HTTPMethodOverrideHandler(r)
	o := handlers.AllowedOrigins([]string{"*"})
	handler = handlers.CORS(o)(r)
	return handler
}
