package main

import "fmt"

const NMAX int = 100

// type tanggal struct {
// 	hari, bulan, tahun int
// }

type pelanggan struct {
	noKontrak, nama, link string
	kapasitas             int
	tglMulai, tglAkhir    string
}

type arrPelanggan [NMAX]pelanggan

func menu(T *arrPelanggan, n *int) {
	var pilihan int

	fmt.Print("\n============ Menu ============\n")
	fmt.Print("     1. Tambah Data\n")
	fmt.Print("     2. Hapus Data\n")
	fmt.Print("     3. Lihat Daftar Pelanggan\n")
	fmt.Print("     4. Cari Data Pelanggan\n")
	fmt.Print("     5. Exit\n\n")

	fmt.Print("Pilih Menu: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		inputData(T, n)
	} else if pilihan == 2 {
		hapusData(T, n)
	} else if pilihan == 3 {
		cetakArray(*T, *n)
	} else if pilihan == 4 {

	} else if pilihan == 5 {

	} else {
		menu(T, n)
	}
}

func inputData(T *arrPelanggan, n *int) {
	var noKontrak string
	fmt.Scan(&noKontrak)
	for noKontrak != "#" {
		if cariIdx(*T, *n, noKontrak, "kontrak") == -1 {
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
	var idx = cariIdx(*T, n, noKontrak, "kontrak")
	fmt.Scan(&T[idx].nama, &T[idx].link, &T[idx].kapasitas, &T[idx].tglMulai, &T[idx].tglAkhir)
}

func hapusData(T *arrPelanggan, n *int) {
	var nama string
	var idx int
	fmt.Print("Nama yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx = cariIdx(*T, *n, nama, "nama")
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	*n--
	menu(T, n)
}

func cariIdx(T arrPelanggan, n int, name string, flag string) int {
	var idx, i int
	var found bool
	i = 0
	idx = -1
	found = false
	if flag == "nama" {
		for i < n && !found {
			if T[i].nama == name {
				idx = i
				found = true
			}
			i++
		}
	} else if flag == "kontrak" {
		for i < n && !found {
			if T[i].noKontrak == name {
				idx = i
				found = true
			}
			i++
		}
	}
	return idx
}

func cetakArray(T arrPelanggan, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], "\n")
	}
	menu(&T, &n)
}

func main() {
	var T arrPelanggan
	var n int
	menu(&T, &n)
}
