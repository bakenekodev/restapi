package main

import (
	"database/sql"
	"log"
	"net/http"
)

// Login struct
type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// CheckPassword function
func CheckPassword(w http.ResponseWriter, r *http.Request) {
	login, ok1 := r.URL.Query()["login"]
	password, ok2 := r.URL.Query()["password"]
	if ok1 && ok2 {
		var pass sql.NullString
		var id sql.NullString
		DB.QueryRow(Queries["selectPassword"], login[0]).Scan(&pass, &id)
		if !pass.Valid || pass.String != password[0] {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Write([]byte(id.String))
		}
	}
}

// CreateLogin function
func CreateLogin(w http.ResponseWriter, r *http.Request) {
	login, ok1 := r.URL.Query()["login"]
	password, ok2 := r.URL.Query()["password"]

	if ok1 && ok2 {
		var l sql.NullString
		DB.QueryRow(Queries["checkLogin"], login[0]).Scan(&l)
		log.Println(l)
		if l.Valid {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var id sql.NullString
			DB.QueryRow(Queries["insertLogin"], login[0], password[0]).Scan(&id)
			log.Println(id)
			if id.Valid {
				w.Write([]byte(id.String))
			}
		}
	}
}
