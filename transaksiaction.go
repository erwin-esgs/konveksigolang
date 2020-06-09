package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//LOGINFORM ...
func Transaksiaction(response http.ResponseWriter, request *http.Request) {

	idtransaksi := request.URL.Query().Get("idtransaksi")
	tolak := request.URL.Query().Get("tolak")
	status, _ := strconv.Atoi((request.URL.Query().Get("status")))
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	eer, err := db.Exec(`UPDATE trans SET status = ?  WHERE idtransaksi = ?;`, status, idtransaksi)
	if eer == nil || err != nil {
		panic(err.Error())
	}

	if tolak == "1" {
		path := "img/bukti/" + idtransaksi + ".jpg"
		_ = os.Remove(path)
	}

	http.Redirect(response, request, "/index", http.StatusSeeOther)
}
