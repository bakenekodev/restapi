package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lib/pq"
)

// Route struct
type Route struct {
	ID       string `json:"id"`
	DriverID string `json:"driver_id"`
}

// CreateDriverRoute adds a driver trip record to database
func CreateDriverRoute(w http.ResponseWriter, r *http.Request) {
	driverID, ok := r.URL.Query()["id"]
	if ok {
		var trip [][]float64
		_ = json.NewDecoder(r.Body).Decode(&trip)

		_, err = DB.Exec(Queries["upsetDriverRoute"], driverID, pq.Array(trip))
		if err != nil {
			panic(err.Error())
		}
		log.Println(driverID)
		log.Println(trip)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// FinishRoute function
func FinishRoute(w http.ResponseWriter, r *http.Request) {
	driverID, ok := r.URL.Query()["id"]
	if ok {
		DB.QueryRow(Queries["deleteRoute"], driverID)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
