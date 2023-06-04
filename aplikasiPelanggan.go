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
	fmt.Print("     2. Ubah Data\n")
	fmt.Print("     3. Hapus Data\n")
	fmt.Print("     4. Lihat Daftar Pelanggan\n")
	fmt.Print("     5. Cari Data Pelanggan\n")
	fmt.Print("     6. Exit\n\n")

	fmt.Scan(&pilihan)
	if pilihan == 1 {
		inputData(T, n)
	} else if pilihan == 2 {

	} else if pilihan == 3 {
		hapusData(T, n)
	} else if pilihan == 4 {
		cetakArray(*T, *n)
	} else if pilihan == 5 {

	} else if pilihan == 6 {

	} else {
		menu(T, n)
	}
}

func inputData(T *arrPelanggan, n *int) {
	var data pelanggan
	fmt.Scan(&data.noKontrak)
	for data.noKontrak != "#" && *n < NMAX {
		fmt.Scan(&data.nama, &data.link, &data.kapasitas, &data.tglMulai, &data.tglAkhir)
		T[*n].nama = data.nama
		T[*n].link = data.link
		T[*n].kapasitas = data.kapasitas
		T[*n].tglMulai = data.tglMulai
		T[*n].tglAkhir = data.tglAkhir
		*n++
		fmt.Scan(&data.noKontrak)
	}
	fmt.Print("SELESAI\n\n")
	menu(T, n)
}

func hapusData(T *arrPelanggan, n *int) {
	var nama string
	var idx int
	fmt.Print("Nama yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx = cariIdxNama(*T, *n, nama)
	for i := idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	*n--
	menu(T, n)
}

func cariIdxNama(T arrPelanggan, n int, name string) int {
	var idx, i int
	var found bool
	idx = -1
	found = false
	for i < n || !found {
		if T[i].nama == name {
			idx = i
			found = true
		}
		i++
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
