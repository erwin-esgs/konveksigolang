package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

//Keranjang ...
func Keranjang(response http.ResponseWriter, request *http.Request) {
	//encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	//decoded, err := base64.StdEncoding.DecodeString(encoded)

	type data struct {
		Username     string
		Isikeranjang []Barang
		Totalharga   int
	}
	//var jsonString = `{"Idproduk": "john wick", "Jumlah": 27 , "Deskripsi": " wick"}`
	//var jsonData = []byte(jsonString)
	var Data data

	session, _ := store.Get(request, "mysession")
	Data.Username = session.Values["usename"].(string)

	cookie, err := request.Cookie("Isikeranjang")
	if err == nil {
		//fmt.Println(request)
		decoded, _ := base64.StdEncoding.DecodeString(cookie.Value)
		//fmt.Println(string(decoded))

		var err = json.Unmarshal([]byte(string(decoded)), &Data.Isikeranjang)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {

			db, err := connect()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer db.Close()
			Data.Totalharga = 0
			for i := 0; i < len(Data.Isikeranjang); i++ {
				rows, err := db.Query("SELECT namaproduk , harga FROM produk WHERE idproduk = ?", Data.Isikeranjang[i].Idproduk)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				for rows.Next() {
					var err = rows.Scan(&Data.Isikeranjang[i].Namaproduk, &Data.Isikeranjang[i].Harga)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}
				Data.Totalharga = Data.Totalharga + (Data.Isikeranjang[i].Harga * Data.Isikeranjang[i].Jumlah)
			}
			//fmt.Println(Data)
		}
	}

	tmp, _ := template.ParseFiles("keranjang.gohtml", "navbar.gohtml")
	tmp.Execute(response, Data)
}
