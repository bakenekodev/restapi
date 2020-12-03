package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Car Struct
type Car struct {
	ID    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Seats string `json:"seats"`
}

// GetCars function
func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cars []Car

	results, err := DB.Query(Queries["selectCars"])
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	for results.Next() {
		var car Car
		err := results.Scan(&car.ID, &car.Mark, &car.Model, &car.Year, &car.Seats)
		if err != nil {
			panic(err.Error())
		}
		cars = append(cars, car)
	}

	err = json.NewEncoder(w).Encode(cars)
	if err != nil {
		panic(err.Error())
	}
}

// GetCar function
func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Test if car id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectCarID"], params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	var car Car

	DB.QueryRow(Queries["selectCarByID"], params["id"]).Scan(&car.ID, &car.Mark, &car.Model, &car.Year, &car.Seats)

	json.NewEncoder(w).Encode(car)
}

// CreateCar function
func CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)

	err := DB.QueryRow(Queries["insertCar"], car.Mark, car.Model, car.Year, car.Seats).Scan(&car.ID)
	if err != nil {
		panic(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(car)
	}
}

// UpdateCar function
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Test if car id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectCarID"], params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)

	_, err := DB.Exec(Queries["updateCarByID"], car.Mark, car.Model, car.Year, car.Seats, params["id"])
	if err != nil {
		panic(err.Error())
	} else {
		w.Write([]byte("resource updated"))
	}
}

// DeleteCar function
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Test if car id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectCarID"], params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	_, err := DB.Exec(Queries["deleteCarByID"], params["id"])
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("resource deleted"))
}
