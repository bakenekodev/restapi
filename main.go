package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	goyesql "github.com/nleof/goyesql"
	"golang.org/x/oauth2/google"
)

// DB is the global MySQL database
var DB *sql.DB

// Queries is a map that stores all the SQL queries
var Queries goyesql.Queries
var err error

func main() {

	// Prepare the database
	//setUpDB()
	DB, err = sql.Open("mysql", "root:ht7n0*h66+@tcp(127.0.0.1:3306)/app")
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

	// Driver
	r.HandleFunc("/api/drivers", CreateDriverRoute).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8880",
	}
	log.Fatal(srv.ListenAndServe())
}

// setUpDB connects to the Cloud SQL throgh a proxy server
func setUpDB() {
	credsFile := "credentials.json"
	SQLScope := "https://www.googleapis.com/auth/sqlservice.admin"
	ctx := context.Background()

	creds, err := ioutil.ReadFile(credsFile)
	if err != nil {
		panic(err.Error())
	}

	cfg, err := google.JWTConfigFromJSON(creds, SQLScope)
	if err != nil {
		panic(err.Error())
	}

	client := cfg.Client(ctx)
	proxy.Init(client, nil, nil)

	cf := mysql.Cfg("sacred-tenure-294609:europe-west3:database", "root", "jyi6hk")
	cf.DBName = "app"
	DB, err = mysql.DialCfg(cf)
	if err != nil {
		panic(err.Error())
	}
}
