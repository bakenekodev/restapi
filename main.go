package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	goyesql "github.com/nleof/goyesql"
)

// DB is the global Postgres database
var DB *sql.DB

// Conection info
const (
	host     = "ec2-54-228-170-125.eu-west-1.compute.amazonaws.com"
	port     = 5432
	user     = "gctgugggvhltaw"
	password = "f832bb546fa4574d3b49c28da0e7c8a007154de81aeaa82a14b62836fa11e929"
	dbname   = "dgobk4k3st1m1"
)

// Queries is a map that stores all the SQL queries
var Queries goyesql.Queries
var err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	// Prepare the database
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected!")

	// Prepare the queries
	Queries = goyesql.MustParseFile("queries.sql")

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	// Login
	r.HandleFunc("/api/login", CheckPassword).Methods("GET")
	r.HandleFunc("/api/login", CreateLogin).Methods("POST")
	r.HandleFunc("/api/login", ChangePassword).Methods("PATCH")

	// User
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/users/{id}", UpdatePos).Methods("PATCH")

	// Car
	r.HandleFunc("/api/cars", GetCars).Methods("GET")
	r.HandleFunc("/api/cars", CreateCar).Methods("POST")
	r.HandleFunc("/api/cars/{id}", GetCar).Methods("GET")
	r.HandleFunc("/api/cars/{id}", UpdateCar).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", DeleteCar).Methods("DELETE")

	// Passenger
	r.HandleFunc("/api/passengers", AcceptDriver).Methods("GET")
	r.HandleFunc("/api/passengers", FindRoute).Methods("POST")
	r.HandleFunc("/api/passengers", DeclineDriver).Methods("DELETE")

	// Driver
	r.HandleFunc("/api/drivers", CheckPassengers).Methods("GET")
	r.HandleFunc("/api/drivers", CreateRoute).Methods("POST")
	r.HandleFunc("/api/drivers", FinishRoute).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println(port)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + port,
	}
	log.Fatal(srv.ListenAndServe())
}
