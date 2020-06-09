package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//ADD PRODUCT
func Tambahproduk(response http.ResponseWriter, request *http.Request) {
	t := time.Now()
	idproduk := string(t.Format("060102150405"))
	//str := string(timeformat) + "aaa"
	//idproduk := strings.Replace(str, "a", "", 3)
	var status int8 = 1
	var gambar string = ""

	if err := request.ParseMultipartForm(1024); err == nil {

		request.ParseForm()
		namaproduk := request.Form.Get("namaproduk")
		harga := request.Form.Get("harga")
		deskripsi := request.Form.Get("deskripsi")

		if request.Method != "POST" {
			http.Error(response, "", http.StatusBadRequest)
			return
		}

		uploadedFile, _, err := request.FormFile("image")
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()
		dir, err := os.Getwd()
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := fmt.Sprintf("%s%s", idproduk, ".jpg") //filepath.Ext(handler.Filename)

		fileLocation := filepath.Join(dir, "img", filename)
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

		db, err := connect()
		defer db.Close()
		rows, err := db.Prepare("INSERT INTO produk(idproduk , namaproduk , harga , deskripsi , status , gambar) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		rows.Exec(idproduk, namaproduk, harga, deskripsi, status, gambar)
		http.Redirect(response, request, "/index", http.StatusSeeOther)

	} else {
		http.Redirect(response, request, "/formtambahproduk", http.StatusSeeOther)
	}
}
