package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User Struct (Model)
type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surmane  string `json:"surname"`
	Phone    string `json:"phone"`
	carID    sql.NullInt64
	Car      *Car `json:"car"`
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
		err := results.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if err != nil {
			panic(err.Error())
		}
		if user.carID.Valid {
			user.Car = new(Car)
			DB.QueryRow(Queries["selectCarByID"], user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
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
		err := result.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Surmane, &user.Phone, &user.carID)
		if err != nil {
			panic(err.Error())
		}
		if user.carID.Valid {
			user.Car = new(Car)
			DB.QueryRow(Queries["selectCarByID"], user.carID).Scan(&user.Car.ID, &user.Car.Mark, &user.Car.Model, &user.Car.Year, &user.Car.Seats)
		}
	}
	json.NewEncoder(w).Encode(user)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Car != nil {
		err := DB.QueryRow(Queries["insertCar"], user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats).Scan(&user.carID)
		if err != nil {
			panic(err.Error())
		}
	}

	err := DB.QueryRow(Queries["insertUser"], user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID).Scan(&user.ID)
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

	DB.QueryRow(Queries["selectCarIdByUser"], user.ID).Scan(&user.carID)

	if user.carID.Valid && user.Car != nil {
		_, err := DB.Exec(Queries["updateCarByID"], user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats, user.carID)
		if err != nil {
			panic(err.Error())
		}
	} else if !user.carID.Valid && user.Car != nil {
		result, err := DB.Exec(Queries["insertCar"], user.Car.Mark, user.Car.Model, user.Car.Year, user.Car.Seats)
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

	_, err = DB.Exec(Queries["updateUserByID"], user.Login, user.Password, user.Name, user.Surmane, user.Phone, user.carID, user.ID)
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
	var carID sql.NullInt64

	DB.QueryRow(Queries["selectCarIdByUser"], userID).Scan(&carID)

	_, err := DB.Exec(Queries["deleteUserByID"], userID)
	if err != nil {
		panic(err.Error())
	}

	if carID.Valid {
		_, err := DB.Exec(Queries["deleteCarByID"], carID)
		if err != nil {
			panic(err.Error())
		}
	}
}

// GetLogin function
func GetLogin(w http.ResponseWriter, r *http.Request) {

	login, ok1 := r.URL.Query()["login"]
	password, ok2 := r.URL.Query()["password"]
	if ok1 && ok2 {
		var pass sql.NullString
		var id sql.NullString
		DB.QueryRow(Queries["selectPassword"], login[0]).Scan(&pass, &id)
		if !pass.Valid || pass.String != password[0] {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Login Failed"))
		} else {
			w.Write([]byte(id.String))
		}

	}
}
