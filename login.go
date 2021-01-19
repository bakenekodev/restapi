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

// ChangePassword function
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	id, ok1 := r.URL.Query()["id"]
	password, ok2 := r.URL.Query()["password"]
	newPass, ok3 := r.URL.Query()["new"]

	if ok1 && ok2 && ok3 {
		var pass sql.NullString
		DB.QueryRow(Queries["selectPasswordById"], id[0]).Scan(&pass)
		if pass.Valid && pass.String == password[0] {
			DB.QueryRow(Queries["updatePassword"], id[0], newPass[0])
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
