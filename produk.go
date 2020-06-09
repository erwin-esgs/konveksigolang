package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Produk(response http.ResponseWriter, request *http.Request) {

	type Data struct {
		Username      string
		Detailproduks Detailproduk
		Status        int
	}

	var data Data
	var detailproduk Detailproduk

	session, _ := store.Get(request, "mysession")
	data.Username = session.Values["usename"].(string)
	data.Status = session.Values["status"].(int)
	detailproduk.Idproduk = request.URL.Query().Get("idproduk")

	//fmt.Println(reflect.TypeOf(idproduk))
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT namaproduk , deskripsi , harga, status FROM produk WHERE idproduk = " + detailproduk.Idproduk)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {

		var err = rows.Scan(&detailproduk.Namaproduk, &detailproduk.Deskripsi, &detailproduk.Harga, &detailproduk.Status)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		detailproduk.Filename = detailproduk.Idproduk + ".jpg"

	}
	data.Detailproduks = detailproduk
	tmp, _ := template.ParseFiles("produk.gohtml", "navbar.gohtml", "tambahkeranjang.gohtml")
	tmp.Execute(response, data)

}
