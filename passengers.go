package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lib/pq"
)

// Passanger struct
type Passanger struct {
	StartLat float64 `json:"start_lat"`
	StartLng float64 `json:"start_lng"`
	EndLat   float64 `json:"end_lat"`
	EndLng   float64 `json:"end_lng"`
	StartR   float64 `json:"start_r"`
	EndR     float64 `json:"end_r"`
}

//FindRoute function
func FindRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var pass Passanger
	err = json.NewDecoder(r.Body).Decode(&pass)
	if err != nil {
		panic(err.Error())
	}

	var d pq.Int64Array
	start := []float64{pass.StartLat, pass.StartLng}
	end := []float64{pass.EndLat, pass.EndLng}
	err = DB.QueryRow(Queries["getDriversFunc"], pq.Array(start), pq.Array(end), pass.StartR).Scan(&d)
	if err != nil {
		panic(err.Error())
	}
	// for _, i := range d {
	// 	var driver User
	// 	err = DB.QueryRow(Queries["selectUserByID"], i).Scan(&driver.ID, &driver.Name, &driver.Surmane, &driver.Phone, &driver.CarID)
	// 	drivers = append(drivers, driver)
	// }

	log.Println(d)
	json.NewEncoder(w).Encode(d)
}

// AcceptDriver function
func AcceptDriver(w http.ResponseWriter, r *http.Request) {
	driverID, ok := r.URL.Query()["driver_id"]
	id, ok2 := r.URL.Query()["id"]
	if ok && ok2 {
		_, err := DB.Exec(Queries["acceptDriver"], id[0], driverID[0])
		if err != nil {
			panic(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeclineDriver func
func DeclineDriver(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if ok {
		_, err := DB.Exec(Queries["declineDriver"], id[0])
		if err != nil {
			panic(err.Error())
		}
	}
}
