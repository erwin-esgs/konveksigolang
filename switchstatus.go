package main

func switchstatus(status int) (stat string) {
	switch status {
	case 1:
		stat = "Belum Dibayar"
	case 2:
		stat = "Menunggu Verifikasi"
	case 3:
		stat = "Diproses"
	case 4:
		stat = "Siap diambil"
	case 5:
		stat = "Selesai"
	default:
		stat = "Ditolak"
	}
	return
}
