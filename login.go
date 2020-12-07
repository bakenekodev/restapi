package main

import (
	"database/sql"
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
