package handler

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rtravitz/coparkfinder/models"
	"net/http"
)

type Env struct {
	DB models.Datastore
}

func (env *Env) NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/parks", env.ParksIndex).
		Methods("GET")
	r.HandleFunc("/api/v1/activities", env.ActivitiesIndex).
		Methods("GET")
	r.HandleFunc("/api/v1/facilities", env.FacilitiesIndex).
		Methods("GET")
	r.HandleFunc("/api/v1/park", env.ParkShow).
		Methods("GET")

	handler := handlers.HTTPMethodOverrideHandler(r)
	o := handlers.AllowedOrigins([]string{"*"})
	handler = handlers.CORS(o)(r)
	return handler
}
