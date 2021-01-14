package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lib/pq"
)

// Route struct
type Route struct {
	DriverID string      `json:"driver_id"`
	Points   [][]float64 `json:"points"`
}

// CreateRoute adds a driver trip record to database
func CreateRoute(w http.ResponseWriter, r *http.Request) {
	driverID, ok := r.URL.Query()["id"]
	if ok {
		var trip [][]float64
		_ = json.NewDecoder(r.Body).Decode(&trip)

		_, err = DB.Exec(Queries["upsetDriverRoute"], driverID[0], pq.Array(trip))
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
		_, err = DB.Exec(Queries["deleteRoute"], driverID[0])
		if err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
