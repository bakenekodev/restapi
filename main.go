package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Driver Struct (Model)
type User struct {
	ID      string `json:"id"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Surmane string `json:"surname"`
	Phone   string `json:"phone"`
	Auto    *Car   `json:"car"`
}

type Car struct {
	ID    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Seats string `json:"seats"`
}

type Passenger struct {
	ID       string `json:"id"`
	TripID   string `json:"trip_id"`
	UserId   string `json:"user_id"`
	StartLat string `json:"start_lat"`
	StartLng string `json:"start_lng"`
	EndLat   string `json:"end_lat"`
	EndLng   string `json:"end_lng"`
}

type Trip struct {
	ID       string `json:"id"`
	DriverId string `json:"driver_id"`
	StartLat string `json:"start_lat"`
	StartLng string `json:"start_lng"`
	EndLat   string `json:"end_lat"`
	EndLng   string `json:"end_lng"`
}

var users []User
var db *sql.DB // maybe make local

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get a single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through users and find with id
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// Add new user
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content.Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	lastUserId, err := strconv.Atoi(users[len(users)-1].ID)
	if err != nil {
		panic(err.Error())
	}
	user.ID = strconv.Itoa(lastUserId + 1)
	users = append(users, user)
	db.Query("insert into user (id, login, password, name, surname, telephone, car_id) values (" 
																					+ user.ID + ", "
																					+ user.Login + ", "
																					+ user.Name + ", "
																					+user.Surmane + ", "
																					+ user.Phone + ", "
																					+carId
																				)
	json.NewEncoder(w).Encode(user)

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Database
	db, err := sql.Open("mysql", "root:ON%6RJ@@tcp(localhost:3306)/schema")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, login, name, surname, telephone, car_id From user")
	defer results.Close()
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		var carId sql.NullString
		err = results.Scan(&user.ID, &user.Login, &user.Name, &user.Surmane, &user.Phone, &carId)
		if err != nil {
			panic(err.Error())
		}
		if carId.Valid {
			user.Auto = new(Car)
			db.QueryRow("SELECT id, mark, model, year, seats FROM car WHERE id = "+carId.String).Scan(&user.Auto.ID, &user.Auto.Mark, &user.Auto.Model, &user.Auto.Year, &user.Auto.Seats)
		}
		users = append(users, user)
	}

	results.Close()
	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/drivers", addUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
