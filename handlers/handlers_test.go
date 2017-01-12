package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivitiesIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/activities", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := NewHandler(tdb)
	handler := http.HandlerFunc(h.ActivitiesIndex)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
