package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Transaksikonfirmasi(response http.ResponseWriter, request *http.Request) {

	if err := request.ParseMultipartForm(1024); err == nil {

		request.ParseForm()
		if request.Method != "POST" {
			http.Error(response, "", http.StatusBadRequest)
			return
		}
		idtransaksi := request.URL.Query().Get("idtransaksi")
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
		filename := fmt.Sprintf("%s%s", idtransaksi, ".jpg") //filepath.Ext(handler.Filename)

		fileLocation := filepath.Join(dir, "img/bukti", filename)
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
		eer, err := db.Exec(`UPDATE trans SET status = ?  WHERE idtransaksi = ?;`, 2, idtransaksi)
		if eer == nil || err != nil {
			panic(err.Error())
		}
		http.Redirect(response, request, "/index", http.StatusSeeOther)

	}
}
