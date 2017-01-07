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

	handler := handlers.HTTPMethodOverrideHandler(r)
	o := handlers.AllowedOrigins([]string{"*"})
	handler = handlers.CORS(o)(r)
	return handler
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
