package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
)

func Registersave(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	username := request.Form.Get("username")
	nama := request.Form.Get("nama")
	b := md5.Sum([]byte(request.Form.Get("password")))
	password := hex.EncodeToString(b[:])
	telepon := request.Form.Get("telepon")
	alamat := request.Form.Get("alamat")

	if request.Method != "POST" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	db, err := connect()
	defer db.Close()
	err = db.QueryRow("SELECT username FROM user WHERE username = ?", username).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		{
			rows, err := db.Prepare("INSERT INTO user(username , password , nama , telepon , alamat , status) VALUES(?,?,?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			rows.Exec(username, password, nama, telepon, alamat, 1)
			http.Redirect(response, request, "/index", http.StatusSeeOther)
		}
	case err != nil:
		fmt.Println("error")
	default:
		http.Redirect(response, request, "/register?username=1", http.StatusSeeOther)
	}

}
