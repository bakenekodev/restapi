package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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

// GetUsers function
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	results, err := DB.Query("SELECT id, login, password, name, surname, telephone, car_id From users")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if err != nil {
			panic(err.Error())
		}
		if user.carID.Valid {
			user.Car = new(Car)
			DB.QueryRow("SELECT id, mark, model, year, seats FROM cars WHERE id = ?", user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
		}
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err.Error())
	}

}

// GetUser function
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow("SELECT id FROM users WHERE id = ? LIMIT 1", params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	result, err := DB.Query("SELECT id, login, password, name, surname, telephone, car_id From users WHERE ID = ?", params["id"])
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
			DB.QueryRow("SELECT id, mark, model, year, seats FROM cars WHERE id = ?", user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
		}
	}
	json.NewEncoder(w).Encode(user)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) { // TODO: make driving_lic and photo base64
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	log.Println(user.Car)
	if user.Car != nil {
		result, err := DB.Exec("INSERT INTO cars(mark, model, year, seats, driving_lic) VALUES (?, ?, ?, ?, 1)", user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats)
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
	_, err = DB.Exec("INSERT INTO users(login, password, name, surname, telephone, photo, car_id) VALUES(?, ?, ?, ?, ?, 1, ?)", user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID)
	if err != nil {
		panic(err.Error())
	}
}

// UpdateUser by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow("SELECT id FROM users WHERE id = ? LIMIT 1", params["id"]).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}

	var user User
	user.ID = params["id"]
	_ = json.NewDecoder(r.Body).Decode(&user)

	DB.QueryRow("SELECT car_id FROM users WHERE id = ?", user.ID).Scan(&user.carID)

	if user.carID.Valid && user.Car != nil {
		_, err := DB.Exec("UPDATE cars set mark = ?, model = ?, year = ?, seats = ?, driving_lic = 1 WHERE id = ?",
			user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats, user.carID)
		if err != nil {
			panic(err.Error())
		}
	} else if user.carID.Valid && user.Car == nil {
		// TODO remove car
	} else if !user.carID.Valid && user.Car != nil {
		result, err := DB.Exec("INSERT INTO cars(mark, model, year, seats, driving_lic) VALUES (?, ?, ?, ?, 1)", user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats)
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

	_, err = DB.Exec("update users set login = ?, password = ?, name = ?, surname = ?, telephone = ?, photo = 1, car_id = ? where id = ?", user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID, user.ID)
	if err != nil {
		panic(err.Error())
	}

}

// DeleteUser by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["id"]

	// Test if user id exists in database
	var testID sql.NullString
	DB.QueryRow("SELECT id FROM users WHERE id = ? LIMIT 1", userID).Scan(&testID)
	if !testID.Valid {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - ID not found"))
		return
	}
	var carID sql.NullInt64

	log.Println(DB.QueryRow("SELECT car_id FROM users WHERE id = ?", userID).Scan(&carID))

	_, err := DB.Exec("delete from users where id = ?", userID)
	if err != nil {
		panic(err.Error())
	}

	if carID.Valid {
		_, err := DB.Exec("delete from cars where id = ?", carID)
		if err != nil {
			panic(err.Error())
		}
	}
}
