package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

const NMAX int = 100

type pelanggan struct {
	noKontrak, nama, link string
	kapasitas             int
	tglMulai, tglAkhir    string
}

type arrPelanggan [NMAX]pelanggan

func header() {
	fmt.Print("┌─┐┌─┐┬  ┬┬┌─┌─┐┌─┐┬    ┌─┐┌─┐┬  ┌─┐┌┐┌┌─┐┌─┐┌─┐┌┐┌\n")
	fmt.Print("├─┤├─┘│  │├┴┐├─┤└─┐│    ├─┘├┤ │  ├─┤││││ ┬│ ┬├─┤│││\n")
	fmt.Print("┴ ┴┴  ┴─┘┴┴ ┴┴ ┴└─┘┴    ┴  └─┘┴─┘┴ ┴┘└┘└─┘└─┘┴ ┴┘└┘")
}

func menu(T *arrPelanggan, n *int) {
	var pilihan int

	fmt.Print("\n====================== MENU ======================\n")
	fmt.Print("                 1. Tambah Data\n")
	fmt.Print("                 2. Hapus Data\n")
	fmt.Print("                 3. Urutkan Data\n")
	fmt.Print("                 4. Lihat Daftar Pelanggan\n")
	fmt.Print("                 5. Cari Data Pelanggan\n")
	fmt.Print("                 6. Exit\n\n")

	fmt.Print("Pilih Menu: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		inputData(T, n)
	} else if pilihan == 2 {
		menuHapus(T, *n)
	} else if pilihan == 3 {
		menuUrut(T, *n)
	} else if pilihan == 4 {
		cetakSemuaArray(*T, *n)
	} else if pilihan == 5 {
		menuCari(*T, *n)
	} else if pilihan == 6 {

	} else {
		menu(T, n)
	}
}

func inputData(T *arrPelanggan, n *int) {
	var noKontrak string
	fmt.Scan(&noKontrak)
	for noKontrak != "#" {
		if cariIdxSeq(*T, *n, noKontrak, "kontrak") == -1 {
			tambahData(T, n, noKontrak)
		} else {
			ubahData(T, *n, noKontrak)
		}
		fmt.Scan(&noKontrak)
	}
	fmt.Print("INPUT SELESAI\n")
	menu(T, n)
}

func tambahData(T *arrPelanggan, n *int, noKontrak string) {
	T[*n].noKontrak = noKontrak
	fmt.Scan(&T[*n].nama, &T[*n].link, &T[*n].kapasitas, &T[*n].tglMulai, &T[*n].tglAkhir)
	*n++

}

func ubahData(T *arrPelanggan, n int, noKontrak string) {
	var idx = cariIdxSeq(*T, n, noKontrak, "kontrak")
	fmt.Scan(&T[idx].nama, &T[idx].link, &T[idx].kapasitas, &T[idx].tglMulai, &T[idx].tglAkhir)
}

func menuHapus(T *arrPelanggan, n int) {
	var pilihan int = 0
	var hapus string = ""
	var idx int = -1

	fmt.Print("\n=============== Hapus Berdasarkan ===============\n")
	fmt.Print("                 1. Nomor Kontrak\n")
	fmt.Print("                 2. Nama Pelanggan\n")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		fmt.Print("Nomor kontrak yang ingin dihapus: ")
		fmt.Scan(&hapus)
		idx = cariIdxSeq(*T, n, hapus, "kontrak")
	} else if pilihan == 2 {
		fmt.Print("Nama yang ingin dihapus: ")
		fmt.Scan(&hapus)
		idx = cariIdxSeq(*T, n, hapus, "nama")
	}
	hapusData(T, &n, idx)
	menu(T, &n)
}

func hapusData(T *arrPelanggan, n *int, idx int) {
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	*n--
}

func cariIdxSeq(T arrPelanggan, n int, cari string, flag string) int { //sequential search
	var idx, i int
	var found bool
	i = 0
	idx = -1
	found = false
	if flag == "nama" {
		for i < n && !found {
			if T[i].nama == cari {
				idx = i
				found = true
			}
			i++
		}
	} else if flag == "kontrak" {
		for i < n && !found {
			if T[i].noKontrak == cari {
				idx = i
				found = true
			}
			i++
		}
	}
	return idx
}

func menuUrut(T *arrPelanggan, n int) {
	var pilihan int
	fmt.Print("\n================ Urut Berdasarkan ================\n")
	fmt.Print("                 1. Nomor Kontrak\n")
	fmt.Print("                 2. Nama Pelanggan\n")
	fmt.Print("                 3. Nama Link\n")
	fmt.Print("                 4. Kapasitas\n")
	fmt.Print("                 5. Tanggal Mulai\n")
	fmt.Print("                 6. Tanggal Akhir\n")

	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		sortNoKontrak(T, n)
	} else if pilihan == 2 {
		sortNama(T, n)
	} else if pilihan == 3 {
		sortLink(T, n)
	} else if pilihan == 4 {
		sortKapasitas(T, n)
	} else if pilihan == 5 {
		sortTanggalMulai(T, n)
	} else if pilihan == 6 {
		sortTanggalAkhir(T, n)
	}
	menu(T, &n)
}

func sortKapasitas(T *arrPelanggan, n int) { //mengecil, insertion
	var pass, i int
	var temp pelanggan

	pass = 1
	for pass < n {
		i = pass - 1
		temp = T[pass]
		for i >= 0 && temp.kapasitas > T[i].kapasitas {
			T[i+1] = T[i]
			i--
		}
		T[i+1] = temp
		pass++
	}
}

func sortNoKontrak(T *arrPelanggan, n int) { //membesar, selection
	var pass, i, idx int
	var temp pelanggan

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if T[i].noKontrak < T[idx].noKontrak {
				idx = i
			}
			i++
		}
		temp = T[idx]
		T[idx] = T[pass-1]
		T[pass-1] = temp
		pass++
	}
}

func sortNama(T *arrPelanggan, n int) { //insertion, membesar
	var pass, i int
	var temp pelanggan

	pass = 1
	for pass < n {
		i = pass - 1
		temp = T[pass]
		for i >= 0 && temp.nama < T[i].nama {
			T[i+1] = T[i]
			i--
		}
		T[i+1] = temp
		pass++
	}
}

func sortLink(T *arrPelanggan, n int) { //insertion, membesar
	var pass, i int
	var temp pelanggan

	pass = 1
	for pass < n {
		i = pass - 1
		temp = T[pass]
		for i >= 0 && temp.link < T[i].link {
			T[i+1] = T[i]
			i--
		}
		T[i+1] = temp
		pass++
	}
}

func sortTanggalMulai(T *arrPelanggan, n int) { //insertion, membesar
	var pass, i int
	var temp pelanggan

	pass = 1
	for pass < n {
		i = pass - 1
		temp = T[pass]
		for i >= 0 && balikTanggal(temp, "mulai") < balikTanggal(T[i], "mulai") {
			T[i+1] = T[i]
			i--
		}
		T[i+1] = temp
		pass++
	}
}

func sortTanggalAkhir(T *arrPelanggan, n int) { //insertion, mengecil
	var pass, i int
	var temp pelanggan

	pass = 1
	for pass < n {
		i = pass - 1
		temp = T[pass]
		for i >= 0 && balikTanggal(temp, "akhir") > balikTanggal(T[i], "akhir") {
			T[i+1] = T[i]
			i--
		}
		T[i+1] = temp
		pass++
	}
}

func balikTanggal(P pelanggan, flag string) string {
	var tgl, date string
	if flag == "mulai" {
		tgl = P.tglMulai
	} else if flag == "akhir" {
		tgl = P.tglAkhir
	}

	date = date + string(tgl[6]) + string(tgl[7]) + string(tgl[8]) + string(tgl[9]) + "-"
	date = date + string(tgl[3]) + string(tgl[4]) + "-"
	date = date + string(tgl[0]) + string(tgl[1])
	return date

}

func menuCari(T arrPelanggan, n int) {
	var pilihan, idx int
	var flag, cari string
	idx = -1
	fmt.Print("\n============= Cari Data Berdasarkan =============\n")
	fmt.Print("                 1. Nomor Kontrak\n")
	fmt.Print("                 2. Nama Pelanggan\n")
	fmt.Print("                 3. Nama Link\n")
	fmt.Print("                 4. Kapasitas\n")
	fmt.Print("                 5. Tanggal Mulai\n")
	fmt.Print("                 6. Tanggal Akhir\n")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		flag = "kontrak"
		fmt.Print("Nomor Kontrak: ")
		fmt.Scan(&cari)
		idx = cariIdxSeq(T, n, cari, flag)
	} else if pilihan == 2 {
		flag = "nama"
		fmt.Print("Nama Pelanggan: ")
		fmt.Scan(&cari)
		idx = cariIdxSeq(T, n, cari, flag)
	} else if pilihan == 3 {
		flag = "link"
		fmt.Print("Nama Link: ")
		fmt.Scan(&cari)
		idx = cariIdxBin(&T, n, cari, flag)
	} else if pilihan == 4 {
		flag = "kapasitas"
		fmt.Print("Kapasitas: ")
		fmt.Scan(&cari)
		idx = cariIdxBin(&T, n, cari, flag)
	} else if pilihan == 5 {
		flag = "mulai"
		fmt.Print("Tanggal Mulai: ")
		fmt.Scan(&cari)
		idx = cariIdxBin(&T, n, cari, flag)
	} else if pilihan == 6 {
		flag = "akhir"
		fmt.Print("Tanggal Akhir:")
		fmt.Scan(&cari)
		idx = cariIdxBin(&T, n, cari, flag)
	}
	if idx == -1 {
		fmt.Print("Data Tidak Ditemukan\n")
	} else {
		cetakArray(T, idx)
	}
	menu(&T, &n)
}

func cariIdxBin(T *arrPelanggan, n int, cari, flag string) int {
	var kiri, kanan, mid, idx int
	var ketemu bool = false
	kiri = 0
	kanan = n - 1

	if flag == "link" {
		sortLink(T, n) //membesar
		for kiri <= kanan && !ketemu {
			mid = (kiri + kanan) / 2
			if T[mid].link == cari {
				idx = mid
				ketemu = true
			} else if T[mid].link > cari {
				kanan = mid - 1
			} else if T[mid].link < cari {
				kiri = mid + 1
			}
		}
	} else if flag == "kapasitas" {
		sortKapasitas(T, n) //mengecil
		for kiri <= kanan && !ketemu {
			mid = (kiri + kanan) / 2
			if strconv.Itoa(T[mid].kapasitas) == cari {
				idx = mid
				ketemu = true
			} else if strconv.Itoa(T[mid].kapasitas) < cari {
				kanan = mid - 1
			} else if strconv.Itoa(T[mid].kapasitas) > cari {
				kiri = mid + 1
			}
		}
	} else if flag == "mulai" {
		sortTanggalMulai(T, n) //membesar
		for kiri <= kanan && !ketemu {
			mid = (kiri + kanan) / 2
			if T[mid].tglMulai == cari {
				idx = mid
				ketemu = true
			} else if T[mid].tglMulai > cari {
				kanan = mid - 1
			} else if T[mid].tglMulai < cari {
				kiri = mid + 1
			}
		}
	} else if flag == "akhir" {
		sortTanggalAkhir(T, n) //mengecil
		for kiri <= kanan && !ketemu {
			mid = (kiri + kanan) / 2
			if T[mid].tglAkhir == cari {
				idx = mid
				ketemu = true
			} else if T[mid].tglAkhir < cari {
				kanan = mid - 1
			} else if T[mid].tglAkhir > cari {
				kiri = mid + 1
			}
		}

	}
	return idx
}

func cetakArray(T arrPelanggan, idx int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "No. Kontrak\t Nama\t Link\t Kapasitas\t Tanggal Mulai\t Tanggal Akhir\t")
	fmt.Fprintln(w, T[idx].noKontrak, "\t", T[idx].nama, "\t", T[idx].link, "\t", T[idx].kapasitas, "\t", T[idx].tglMulai, "\t", T[idx].tglAkhir, "\t")
	w.Flush()
	//fmt.Print(T[idx].noKontrak, T[idx].nama, T[idx].link, T[idx].kapasitas, T[idx].tglMulai, T[idx].tglAkhir)
}

func cetakSemuaArray(T arrPelanggan, n int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "No. Kontrak\t Nama\t Link\t Kapasitas\t Tanggal Mulai\t Tanggal Akhir\t")
	for i := 0; i < n; i++ {
		fmt.Fprintln(w, T[i].noKontrak, "\t", T[i].nama, "\t", T[i].link, "\t", T[i].kapasitas, "\t", T[i].tglMulai, "\t", T[i].tglAkhir, "\t")
	}
	w.Flush()
	// fmt.Print("| No Kontrak\t| Nama Pelanggan\t| Link\t| Kapasitas\t| Tanggal Mulai\t| Tanggal Akhir\n")
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("| %v\t| %v\t\t\t| %v\t| %v\t| %v\t| %v\n", T[i].noKontrak, T[i].nama, T[i].link, T[i].kapasitas, T[i].tglMulai, T[i].tglAkhir)
	// }
	menu(&T, &n)
}

func main() {
	var T arrPelanggan
	var n int

	header()
	menu(&T, &n)
}
