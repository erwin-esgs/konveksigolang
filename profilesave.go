package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
)

func Profilesave(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	username := request.Form.Get("username")
	nama := request.Form.Get("nama")
	b := md5.Sum([]byte(request.Form.Get("oldpassword")))
	oldpassword := hex.EncodeToString(b[:])
	b = md5.Sum([]byte(request.Form.Get("password")))
	password := hex.EncodeToString(b[:])
	telepon := request.Form.Get("telepon")
	alamat := request.Form.Get("alamat")

	if request.Method != "POST" {
		http.Error(response, "", http.StatusBadRequest)
		return
	}
	db, err := connect()
	defer db.Close()
	err = db.QueryRow("SELECT username FROM user WHERE username = ? AND password =? ", username, oldpassword).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		http.Redirect(response, request, "/register?username=1", http.StatusSeeOther)
	case err != nil:
		fmt.Println("error")
	default:
		{
			_, err := db.Exec(`UPDATE user SET password = ? , nama = ? , telepon = ? , alamat = ? WHERE username = ?;`, password, nama, telepon, alamat, username)
			if err != nil {
				panic(err.Error())
			}
			http.Redirect(response, request, "/index", http.StatusSeeOther)

		}

	}

}
