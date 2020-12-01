package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	goyesql "github.com/nleof/goyesql"
)

// DB is the global MySQL database
var DB *sql.DB

// Queries is a map that stores all the SQL queries
var Queries goyesql.Queries
var err error

func main() {

	// Prepare the database
	DB, err = sql.Open("mysql", "root:jyi6hk@tcp(34.89.173.199:3306)/app")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()

	// Prepare the queries
	Queries = goyesql.MustParseFile("queries.sql")

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	// User
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	// Car
	r.HandleFunc("/api/cars", GetCars).Methods("GET")
	r.HandleFunc("/api/cars", CreateCar).Methods("POST")
	r.HandleFunc("/api/cars/{id}", GetCar).Methods("GET")
	r.HandleFunc("/api/cars/{id}", UpdateCar).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", DeleteCar).Methods("DELETE")

	// Driver
	r.HandleFunc("/api/drivers", CreateDriverRoute).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	log.Fatal(srv.ListenAndServe())
}
