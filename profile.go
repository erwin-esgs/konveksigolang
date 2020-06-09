package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//Admin ...
func Profile(response http.ResponseWriter, request *http.Request) {

	type Data struct {
		Username string
		Password string
		Nama     string
		Telepon  string
		Alamat   string
	}

	var data Data

	session, _ := store.Get(request, "mysession")
	data.Username = session.Values["usename"].(string)
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT nama, telepon, alamat FROM user WHERE username = ?", data.Username)
	if err != nil {
		return
	}

	for rows.Next() {
		var err = rows.Scan(&data.Nama, &data.Telepon, &data.Alamat)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	}

	tmp, _ := template.ParseFiles("profile.gohtml", "navbar.gohtml")
	tmp.Execute(response, data)
}
