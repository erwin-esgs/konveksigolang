package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Daftarproduk(response http.ResponseWriter, request *http.Request) {

	type Data struct {
		Username string
		Produks  []Barang
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
	var produk Barang

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
		produk.Harga = harga
		produk.Status = status
		produk.Filename = idproduk + ".jpg"
		data.Produks = append(data.Produks, produk)

	}

	tmp, _ := template.ParseFiles("daftarproduk.gohtml", "navbar.gohtml")
	tmp.Execute(response, data)
}
