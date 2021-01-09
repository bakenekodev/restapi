package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User Struct (Model)
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surmane string `json:"surname"`
	Phone   string `json:"phone"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
	CarID   string `json:"car_id"`
	carID   sql.NullString
}

// GetUsers function
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	results, err := DB.Query(Queries["selectUsers"])
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if user.carID.Valid {
			user.CarID = user.carID.String
		}
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err.Error())
	}

}

// GetUser gets the user json by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectUserID"], params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	result, err := DB.Query(Queries["selectUserByID"], params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user User

	for result.Next() {
		err := result.Scan(&user.ID, &user.Name, &user.Surmane, &user.Phone, &user.Lat, &user.Lng, &user.carID)
		if user.carID.Valid {
			user.CarID = user.carID.String
		}
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := DB.QueryRow(Queries["insertUser"], user.ID, user.Name, user.Surmane, user.Phone, user.carID).Scan(&user.ID)
	if err != nil {
		panic(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// UpdateUser by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectUserID"], params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	var user User
	user.ID = params["id"]
	_ = json.NewDecoder(r.Body).Decode(&user)

	_, err = DB.Exec(Queries["updateUserByID"], user.Name, user.Surmane, user.Phone, user.carID, user.ID)
	if err != nil {
		panic(err.Error())
	} else {
		w.Write([]byte("resource updated"))
	}

}

// DeleteUser by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["id"]

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow(Queries["selectUserID"], userID).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	_, err := DB.Exec(Queries["deleteUserByID"], userID)
	if err != nil {
		panic(err.Error())
	}
}
