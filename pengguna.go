package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//Pengguna ...
func Pengguna(response http.ResponseWriter, request *http.Request) {

	type Produk struct {
		Idproduk   string
		Namaproduk string
		Harga      string
		Status     int
		Filename   string
	}
	type Data struct {
		Username string
		Produks  []Produk
	}
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idproduk, namaproduk , harga, status  FROM produk ORDER BY idproduk DESC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var idproduk, namaproduk string
	var harga, status int
	var data Data
	var produk Produk

	session, _ := store.Get(request, "mysession")
	data.Username = session.Values["usename"].(string)

	for rows.Next() {
		var err = rows.Scan(&idproduk, &namaproduk, &harga, &status)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		produk.Idproduk = idproduk
		produk.Namaproduk = namaproduk
		produk.Harga = strconv.Itoa(harga)
		produk.Status = status
		produk.Filename = idproduk + ".jpg"
		data.Produks = append(data.Produks, produk)

	}

	tmp, _ := template.ParseFiles("pengguna.gohtml", "navbar.gohtml", "tambahkeranjang.gohtml")
	tmp.Execute(response, data)
}
