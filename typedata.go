package main

type Trans struct {
	Idtransaksi string
	Totalharga  int
	Status      string
	Statusint   int
}
type Transaksi struct {
	Idproduk   string
	Namaproduk string
	Harga      string
	Jumlah     string
	Keterangan string
}

type Detailproduk struct {
	Idproduk   string
	Namaproduk string
	Deskripsi  string
	Harga      int
	Status     int
	Filename   string
}
type Barang struct {
	Idproduk   string
	Namaproduk string
	Jumlah     int
	Harga      int
	Keterangan string
	Status     int
	Filename   string
}
