package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

//LOGINFORM ...
func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	hash := md5.Sum([]byte(request.Form.Get("password")))
	password := hex.EncodeToString(hash[:])

	session, _ := store.Get(request, "mysession")
	if session.Values["usename"] == nil {
		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()
		rows, err := db.Query("SELECT username , password , status FROM user WHERE username = ?", username)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for rows.Next() {
			var usernamedb, passdb string
			var status int
			var err = rows.Scan(&usernamedb, &passdb, &status)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if username == usernamedb && password == passdb {

				if err != nil {
					http.Error(response, err.Error(), http.StatusInternalServerError)
					return
				}
				session.Values["usename"] = username
				session.Values["status"] = status
				err = session.Save(request, response)
				if err != nil {
					http.Error(response, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Println(session.Values["usename"], "Logged in")
			}
		}
	}
	http.Redirect(response, request, "/index", http.StatusSeeOther)
}
