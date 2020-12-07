package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Route struct
type Route struct {
	ID        string `json:"id"`
	DriverID  string `json:"driver_id"`
	StartLat  string `json:"start_lat"`
	StartLng  string `json:"start_lng"`
	EndLat    string `json:"end_lat"`
	EndLng    string `json:"end_lng"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// CreateDriverRoute adds a driver trip record to database
func CreateDriverRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := r.URL.Query()["id"]
	if ok {
		var trip [][]float32
		_ = json.NewDecoder(r.Body).Decode(&trip)

		log.Println(id)
		log.Println(trip)
	}

	//_, err := DB.Exec(Queries["insertDriverRoute"], trip.DriverID, trip.StartLat, trip.StartLng, trip.EndLat, trip.EndLng, trip.StartTime, trip.EndTime)
	// if err != nil {
	// 	panic(err.Error())
	// }
}
