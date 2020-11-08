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
	"golang.org/x/oauth2/google"
)

// DB - global MySQL database
var DB *sql.DB
var err error

func main() {

	// Prepare the database
	setUpDB()
	defer DB.Close()

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// setUpDB - connects to the Cloud SQL throgh a proxy server
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
