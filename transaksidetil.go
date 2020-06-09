package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

//Keranjang ...
func Transaksidetil(response http.ResponseWriter, request *http.Request) {
	//encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	//decoded, err := base64.StdEncoding.DecodeString(encoded)

	type data struct {
		Username    string
		Idtransaksi string
		Transaksis  []Transaksi
		Totalharga  int
		Statustrans int
		Statusid    int
	}
	var Data data
	session, _ := store.Get(request, "mysession")
	Data.Username = session.Values["usename"].(string)
	Data.Statusid = session.Values["status"].(int)
	Data.Idtransaksi = request.URL.Query().Get("idtransaksi")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT json_str , totalharga , status FROM trans WHERE idtransaksi = ?", Data.Idtransaksi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var decoded string
		var err = rows.Scan(&decoded, &Data.Totalharga, &Data.Statustrans)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = json.Unmarshal([]byte(string(decoded)), &Data.Transaksis)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println(Data.Transaksis)
	}

	tmp, _ := template.ParseFiles("transaksidetil.gohtml", "navbar.gohtml")
	tmp.Execute(response, Data)
}
