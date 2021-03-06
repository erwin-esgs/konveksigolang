package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Penggunatrans(response http.ResponseWriter, request *http.Request) {

	type Data struct {
		Username   string
		Transaksis []Trans
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

	rows, err := db.Query("SELECT idtransaksi, totalharga, status FROM trans WHERE username = ? ORDER BY idtransaksi DESC", data.Username)
	if err != nil {
		return
	}

	for rows.Next() {
		var trans Trans
		var status int
		var err = rows.Scan(&trans.Idtransaksi, &trans.Totalharga, &status)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		trans.Status = switchstatus(status)
		data.Transaksis = append(data.Transaksis, trans)

	}

	tmp, _ := template.ParseFiles("penggunatrans.gohtml", "navbar.gohtml")
	tmp.Execute(response, data)
}
