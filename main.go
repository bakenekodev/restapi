package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// User Struct (Model)
type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"` // TODO: secure authentication
	Name     string `json:"name"`
	Surmane  string `json:"surname"`
	Phone    string `json:"phone"`
	carID    sql.NullInt64
	Car      *Car `json:"car"`
}

// Car Struct
type Car struct {
	ID    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Seats string `json:"seats"`
}

// Passenger struct
type Passenger struct {
	ID       string `json:"id"`
	TripID   string `json:"trip_id"`
	UserID   string `json:"user_id"`
	StartLat string `json:"start_lat"`
	StartLng string `json:"start_lng"`
	EndLat   string `json:"end_lat"`
	EndLng   string `json:"end_lng"`
}

// Trip struct
type Trip struct {
	ID       string `json:"id"`
	DriverID string `json:"driver_id"`
	StartLat string `json:"start_lat"`
	StartLng string `json:"start_lng"`
	EndLat   string `json:"end_lat"`
	EndLng   string `json:"end_lng"`
}

var db *sql.DB
var err error

func main() {
	// Init Router
	r := mux.NewRouter()

	// Database
	db, err = sql.Open("mysql", "root:jyi6hk@tcp(34.89.173.199:3306)/app")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	results, err := db.Query("SELECT id, login, password, name, surname, telephone, car_id From users")
	defer results.Close()
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if err != nil {
			panic(err.Error())
		}
		if user.carID.Valid {
			user.Car = new(Car)
			db.QueryRow("SELECT id, mark, model, year, seats FROM cars WHERE id = ?", &user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Get a single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	result, err := db.Query("SELECT id, login, password, name, surname, telephone, car_id From users WHERE ID = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user User

	for result.Next() {
		err := result.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if err != nil {
			panic(err.Error())
		}
		if user.carID.Valid {
			user.Car = new(Car)
			db.QueryRow("SELECT id, mark, model, year, seats FROM cars WHERE id = ?", &user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
		}
	}
	json.NewEncoder(w).Encode(user)
}

// Create a new user
// TODO: make driving_lic and photo base64
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	log.Println(user.Car)
	if user.Car != nil {
		result, err := db.Exec("INSERT INTO cars(mark, model, year, seats, driving_lic) VALUES (?, ?, ?, ?, 1)", user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats)
		if err != nil {
			panic(err.Error())
		}

		user.carID.Int64, err = result.LastInsertId()
		if err != nil {
			panic(err.Error())
		} else {
			user.carID.Valid = true
		}
	}
	_, err = db.Exec("INSERT INTO users(login, password, name, surname, telephone, photo, car_id) VALUES(?, ?, ?, ?, ?, 1, ?)", user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID)
	if err != nil {
		panic(err.Error())
	}
}

// Update an existing user by id
// TODO: add car if null (add insert if user.Car != nil and user.carID.Valid = true)
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]

	db.QueryRow("SELECT car_id FROM users WHERE id = ?", user.ID).Scan(&user.carID)

	if user.carID.Valid {
		_, err := db.Exec("UPDATE cars set mark = ?, model = ?, year = ?, seats = ?, driving_lic = 1 WHERE id = ?",
			user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats, user.carID)
		if err != nil {
			panic(err.Error())
		}
	}

	_, err = db.Exec("update users set login = ?, password = ?, name = ?, surname = ?, telephone = ?, photo = 1, car_id = ? where id = ?", user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID, user.ID)
	if err != nil {
		panic(err.Error())
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["id"]
	var carID sql.NullInt64

	db.QueryRow("SELECT car_id FROM users WHERE id = ?", userID).Scan(&carID)

	_, err := db.Exec("delete from users where id = ?", userID)
	if err != nil {
		panic(err.Error())
	}

	if carID.Valid {
		_, err := db.Exec("delete from cars where id = ?", carID)
		if err != nil {
			panic(err.Error())
		}
	}
}
