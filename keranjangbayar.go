package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//ADD PRODUCT
func Keranjangbayar(response http.ResponseWriter, request *http.Request) {

	if err := request.ParseMultipartForm(1024); err == nil {

		if request.Method != "POST" {
			http.Error(response, "", http.StatusBadRequest)
			return
		}
		request.ParseForm()

		session, _ := store.Get(request, "mysession")
		username := session.Values["usename"].(string)

		t := time.Now()
		Idtransaksi := string(t.Format("060102150405"))

		var transaksi []Transaksi
		var trans1 Transaksi
		for i, _ := range request.Form["idproduk"] {
			trans1.Idproduk = request.Form["idproduk"][i]
			trans1.Namaproduk = request.Form["namaproduk"][i]
			trans1.Harga = request.Form["harga"][i]
			trans1.Jumlah = request.Form["jumlah"][i]
			trans1.Keterangan = request.Form["keterangan"][i]
			transaksi = append(transaksi, trans1)
		}

		Totalharga := request.Form["totalharga"][0]

		for i, _ := range transaksi { // search file
			var imgname string
			imgname = "image" + string(transaksi[i].Idproduk)
			uploadedFile, _, err := request.FormFile(imgname)
			if err != nil {
				continue
			}
			defer uploadedFile.Close()
			dir, err := os.Getwd()
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
			filename := fmt.Sprintf("%s%s%s%s", Idtransaksi, "req", transaksi[i].Idproduk, ".jpg") //filepath.Ext(handler.Filename)

			fileLocation := filepath.Join(dir, "img/req", filename)
			targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
			defer targetFile.Close()
			if _, err := io.Copy(targetFile, uploadedFile); err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				return
			}
		} // end for search file

		str, err := json.Marshal(transaksi)
		if err != nil {
			return
		}
		json_str := string(str)

		db, err := connect()
		if err != nil {
			return
		}
		defer db.Close()
		rows, err := db.Prepare("INSERT INTO trans (idtransaksi, username, json_str, totalharga,status) VALUES(?,?,?,?,?)")
		if err != nil {
			return
		}
		rows.Exec(Idtransaksi, username, json_str, Totalharga, 1)
		//fmt.Println(json_str)
		http.Redirect(response, request, "/index", http.StatusSeeOther)

	}
}

/*
	images := request.MultipartForm.File["image"] //multi file
	for i, _ := range images {
		img, err := images[i].Open() //open image
		defer img.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		dir, err := os.Getwd() //get work directory
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := fmt.Sprintf("%d%s", i, ".jpg")                                  //filepath.Ext(handler.Filename)  rename
		fileLocation := filepath.Join(dir, "img/req", filename)                     //locate file
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666) //target file to location
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()
		if _, err := io.Copy(targetFile, img); err != nil { //copy to target file
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
	}*/
