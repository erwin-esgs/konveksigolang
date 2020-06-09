package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

// user     => root
// password =>
// host     => 127.0.0.1 atau localhost
// port     => 3306
// dbname   => nama database
func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/konveksi")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Index(response http.ResponseWriter, request *http.Request) {

	session, err := store.Get(request, "mysession")
	//fmt.Println(reflect.TypeOf(session))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["usename"] == nil {
		data := make(map[string]interface{})
		//username, ok := request.URL.Query()["username"]
		username := request.URL.Query().Get("username")
		data["username"] = username

		fmt.Println(data["username"])
		tmp, _ := template.ParseFiles("login.html")
		tmp.Execute(response, data)
	} else {
		if session.Values["status"] == 0 {
			Admin(response, request)
		} else {
			Pengguna(response, request)
		}
	}
}

//Login ...
func login(response http.ResponseWriter, request *http.Request) {
	Login(response, request)
}

//Logout ...
func logout(response http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "mysession")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(session.Values["usename"], ".Logout")
	session.Values["usename"] = nil
	err = session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	time.Sleep(1 * time.Second)
	http.Redirect(response, request, "/index", http.StatusSeeOther)
}

// Register ...
func register(response http.ResponseWriter, request *http.Request) {
	type Data struct {
		Username string
	}
	var data Data
	data.Username = request.URL.Query().Get("username")
	tmp, _ := template.ParseFiles("register.gohtml")
	tmp.Execute(response, data)
}
func registersave(response http.ResponseWriter, request *http.Request) {
	Registersave(response, request)
}

//PROFILE
func profile(response http.ResponseWriter, request *http.Request) {
	Profile(response, request)
}

//PROFILEsave
func profilesave(response http.ResponseWriter, request *http.Request) {
	Profilesave(response, request)
}

// ADDPRODUCT ...
func formtambahproduk(response http.ResponseWriter, request *http.Request) {
	session, nil := store.Get(request, "mysession")
	username := session.Values["usename"].(string)
	if username == "admin" {
		tmp, _ := template.ParseFiles("tambahproduk.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/index", http.StatusSeeOther)
	}
}

//ADD PRODUCT
func tambahproduk(response http.ResponseWriter, request *http.Request) {
	Tambahproduk(response, request)
}

//PRODUCT LIST
func daftarproduk(response http.ResponseWriter, request *http.Request) {
	Daftarproduk(response, request)
}

//PRODUCT
func produk(response http.ResponseWriter, request *http.Request) {
	Produk(response, request)
}

//Keranjang ...
func keranjang(response http.ResponseWriter, request *http.Request) {
	Keranjang(response, request)
}

//Keranjang bayar ...
func keranjangbayar(response http.ResponseWriter, request *http.Request) {
	Keranjangbayar(response, request)
}

//Keranjang bayar ...
func penggunatrans(response http.ResponseWriter, request *http.Request) {
	Penggunatrans(response, request)
}

func transaksidetil(response http.ResponseWriter, request *http.Request) {
	Transaksidetil(response, request)
}
func transaksiaction(response http.ResponseWriter, request *http.Request) {
	Transaksiaction(response, request)
}
func transaksikonfirmasi(response http.ResponseWriter, request *http.Request) {
	Transaksikonfirmasi(response, request)
}
func transaksiverifikasi(response http.ResponseWriter, request *http.Request) {
	Transaksiverifikasi(response, request)
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/index", Index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
	http.HandleFunc("/registersave", registersave)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/profilesave", profilesave)
	http.HandleFunc("/formtambahproduk", formtambahproduk)
	http.HandleFunc("/tambahproduk", tambahproduk)
	http.HandleFunc("/daftarproduk", daftarproduk)
	http.HandleFunc("/produk", produk)
	http.HandleFunc("/keranjang", keranjang)
	http.HandleFunc("/keranjangbayar", keranjangbayar)
	http.HandleFunc("/penggunatrans", penggunatrans)
	http.HandleFunc("/transaksidetil", transaksidetil)
	http.HandleFunc("/transaksiaction", transaksiaction)
	http.HandleFunc("/transaksikonfirmasi", transaksikonfirmasi)
	http.HandleFunc("/transaksiverifikasi", transaksiverifikasi)
	http.ListenAndServe(":3000", nil)
}
