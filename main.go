package main

import "fmt"

type project struct {
	id        string
	nama      string
	tanggal   string // format YYYYMMDD
	teknologi string // Contoh: Python, Tableau, Scikit-Learn, dll
	kategori  string // Contoh: Machine learning, Data visualization, dll
	kesulitan int    // 1-5
}

const NMAX int = 100

type tabProject [NMAX]project

func menu() {
	fmt.Println("==================================")
	fmt.Println("  PORTOFOLIO DATA SCIENCE PEMULA  ")
	fmt.Println("==================================")
	fmt.Println("1. User Projects")
	fmt.Println("2. Cari Project")
	fmt.Println("3. Urutkan Project")
	fmt.Println("4. Tampilkan Keahlian User")
	fmt.Println("5. Tampilkan Statistik Project")
	fmt.Println("6. Keluar")
	fmt.Println("==================================")
}

func main() {
	var pilihan int
	var data tabProject
	var nProject int

	for {
		menu()
		fmt.Println("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		if pilihan == 6 {
			fmt.Println("Keluar dari program...")
			return
		}

		switch pilihan {
		case 1:
			var pilih int

			pilih = userProject(data, nProject)

			switch pilih {
			case 1:
				tampilkanProject(data, nProject)
			case 2:
				tambah(&data, &nProject)
			case 3:
				ubah(&data, nProject)
			case 4:
				hapus(&data, &nProject)
			case 5:
				fmt.Println("Kembali ke menu utama...")
			default:
				fmt.Println("Pilihan tidak valid!")
			}

		case 2:
			fmt.Println("==================================")
			fmt.Println("-> 1. Cari berdasarkan kategori")
			fmt.Println("-> 2. Cari berdasarkan tingkat kesulitan")
			fmt.Println("-> 3. Kembali")
			fmt.Println("==================================")

			var pilih int
			fmt.Scan(&pilih)

			switch pilih {
			case 1:
				cariKategori()
			case 2:
				cariTingkatKesulitan()
			case 3:
				fmt.Println("Kembali ke menu...")
			default:
				fmt.Println("Pilihan tidak valid!")
			}

		case 3:
			fmt.Println("==================================")
			fmt.Println("-> 1. Urutkan berdasarkan tanggal")
			fmt.Println("-> 2. Urutkan berdasarkan tingkat kesulitan")
			fmt.Println("-> 3. Kembali")
			fmt.Println("==================================")

			var pilih int
			fmt.Scan(&pilih)

			switch pilih {
			case 1:
				urutkanTanggal()
			case 2:
				urutkanTingkatKesulitan()
			case 3:
				fmt.Println("Kembali ke menu...")
			default:
				fmt.Println("Pilihan tidak valid!")
			}

		case 4:
			tampilkanKeahlian(data, nProject)
		case 5:
			tampilkanStatistik(data, nProject)
		default:
			fmt.Println("Harap masukkan angka yang valid (1-6)")
		}
	}
}

/*
	Prosedur untuk menampilkan menu lanjutan dari 1. User Projects.
	Parameter:
	- P: Daftar tabel semua project yang ada.
	- nP: Jumlah project saat ini
*/
func userProject(P tabProject, nP int) int {
	var pilih int

	fmt.Println("==================================")
	fmt.Println("-> 1. Tampilkan Semua Project")
	fmt.Println("-> 2. Tambah Project")
	fmt.Println("-> 3. Ubah Project")
	fmt.Println("-> 4. Hapus Project")
	fmt.Println("-> 5. Kembali")
	fmt.Println("==================================")

	fmt.Print("Pilih menu (1-5): ")
	fmt.Scan(&pilih)

	return pilih
}

/*
	Prosedur untuk menampilkan semua daftar project.
	Parameter:
	- P: Daftar tabel semua project yang ada.
	- nP: Jumlah project saat ini
*/
func tampilkanProject(P tabProject, nP int) {
	if nP == 0 { // Menampilkan pesan jika belum ada project yang tersedia di tabel
		fmt.Println("Data project belum tersedia, harap masukkan data project.")
		return // Berhenti karena tidak ada untuk dicetak
	}

	fmt.Println("=========================================================================================")
	fmt.Println("                                DAFTAR PROJECT DATA SCIENCE                              ")
	fmt.Println("=========================================================================================")
	fmt.Printf("| %-10s | %-20s | %-10s | %-15s | %-15s |%-10s |\n",
		"ID Project", "Nama", "Tanggal", "Teknologi", "Kategori", "Tingkat Kesulitan")

	fmt.Println("-----------------------------------------------------------------------------------------")

	for i := 0; i < nP; i++ { // Menampilkan daftar project dengan perulangan.
		fmt.Printf("| %-10s | %-20s | %-10s | %-15s | %-15s |%-10d |\n", P[i].id, P[i].nama, P[i].tanggal, P[i].teknologi, P[i].kategori, P[i].kesulitan)
	}
	fmt.Println("=========================================================================================")
}

/*
	Prosedur untuk menambah project ke daftar.
	Parameter:
	- P: Pointer ke array tabProject, memungkinkan modifikasi data di main
	- nP: Pointer ke jumlah project saat ini, memungkinkan pembaruan jumlah project
*/
func tambah(P *tabProject, nP *int) {
	// Membuat ID otomatis, contoh: PRJ-1, PRJ-2, dst.
	P[*nP].id = fmt.Sprintf("PRJ-%d", *nP+1)
	fmt.Printf("ID project: %s\n", P[*nP].id)

	fmt.Print("Nama Project: ")
	fmt.Scan(&P[*nP].nama)
	fmt.Print("Tanggal Project (YYYYMMDD): ")
	fmt.Scan(&P[*nP].tanggal)
	fmt.Print("Teknologi Project: ")
	fmt.Scan(&P[*nP].teknologi)
	fmt.Print("Kategori Project: ")
	fmt.Scan(&P[*nP].kategori)
	fmt.Print("Tingkat Kesulitan Project (1-5): ")
	fmt.Scan(&P[*nP].kesulitan)

	*nP++ // Menambah jumlah project setelah penambahan.
	fmt.Println("Project berhasil ditambahkan!")
	fmt.Println()
}

/*
	Prosedur untuk mengubah/mengedit project dari daftar berdasarkan nomor.
	Parameter:
	- P: Pointer ke array tabProject, memungkinkan modifikasi data di main.
	- nP: Jumlah project saat ini.
*/
func ubah(P *tabProject, nP int) {
	tampilkanProject(*P, nP)
	fmt.Print("Masukkan ID project yang ingin diubah (contoh: PRJ-1): ")

	var targetID string
	var idxKetemu int
	fmt.Scan(&targetID) // Membaca nomor project yang ingin diubah

	// Nanti pakai Sequential Search untuk mencari indeksnya ! adel
	idxKetemu = seqSearch(*P, nP, targetID)

	if idxKetemu != -1 {
		fmt.Print("Nama Project baru: ")
		fmt.Scan(&P[idxKetemu].nama) // Mengubah nama
		fmt.Print("Tanggal baru (YYYYMMDD): ")
		fmt.Scan(&P[idxKetemu].tanggal) // Mengubah tanggal
		fmt.Print("Nama Teknologi baru: ")
		fmt.Scan(&P[idxKetemu].teknologi) // Mengubah teknologi
		fmt.Print("Nama Kategori baru: ")
		fmt.Scan(&P[idxKetemu].kategori) // Mengubah kategori
		fmt.Print("Nama Kesulitan baru (1-5): ")
		fmt.Scan(&P[idxKetemu].kesulitan) // Mengubah kesulitan

		fmt.Println()
		fmt.Println("Project berhasil diubah!")
	} else {
		fmt.Println("Nomor Project tidak valid!")
		fmt.Println()
	}
	fmt.Println()
}

/*
	Prosedur untuk menghapus project dari daftar berdasarkan nomor.
	Parameter:
	- P: Pointer ke array tabProject, memungkinkan modifikasi data di main
	- nP: Pointer ke jumlah project saat ini, memungkinkan pembaruan jumlah project
*/
func hapus(P *tabProject, nP *int) {
	tampilkanProject(*P, *nP) // Menampilkan daftar project untuk dipilih
	fmt.Print("Masukkan ID project yang ingin dihapus (contoh: PRJ-1): ")

	var targetID string
	var idxKetemu int
	fmt.Scan(&targetID) // Membaca nomor project yang ingin diubah

	// Nanti pakai Sequential Search untuk mencari indeksnya
	idxKetemu = seqSearch(*P, *nP, targetID)

	if idxKetemu >= 0 && idxKetemu < *nP {
		for i := idxKetemu; i < *nP-1; i++ {
			P[i] = P[i+1]
		}
		*nP-- // Mengurangi jumlah project setelah satu nilai dihapus.
		fmt.Printf("Project nomor %d berhasil dihapus.\n", idxKetemu+1)
	} else {
		fmt.Println("Nomor Project tidak valid.") // Pesan ketika indeks yang ingin dihapus tidak valid.
	}
	fmt.Println()
}

/*
	Prosedur untuk menampilkan semua keahlian user.
	Parameter:
	- P: Daftar tabel semua project yang ada.
	- nP: Jumlah project saat ini
*/
func tampilkanKeahlian(P tabProject, nP int) {
	if nP == 0 {
		fmt.Println("Belum ada project. Tambahkan data project terlebih dahulu.")
		return
	}

	var arrKeahlian [NMAX]string
	var nKeahlian int

	for i := 0; i < nP; i++ {
		var ditemukan bool = false

		// Cek apakah teknologi pada P[i] sudah tersimpan di arrKeahlian
		for j := 0; j < nKeahlian; j++ {
			if arrKeahlian[j] == P[i].teknologi {
				ditemukan = true // tandai jika sudah ada
				break            // hentikan karena sudah ketemu
			}
		}

		if !ditemukan {
			arrKeahlian[nKeahlian] = P[i].teknologi
			nKeahlian++
		}
	}
	// Cetak keahlian user
	fmt.Println("==================================")
	fmt.Println("       DAFTAR KEAHLIAN USER       ")
	fmt.Println("==================================")
	fmt.Println("Berdasarkan project yang dikerjakan, Anda telah menguasai:")
	for i := 0; i < nKeahlian; i++ {
		fmt.Printf("%d. %s\n", i+1, arrKeahlian[i])
	}
	fmt.Println("==================================")
	fmt.Println()
}

/*
	Prosedur untuk menampilkan statistik proyek yang sudah dikerjakan user.
	Parameter:
	- P: Daftar tabel semua project yang ada.
	- nP: Jumlah project saat ini
*/
func tampilkanStatistik(P tabProject, nP int) {
	if nP == 0 {
		fmt.Println("Belum ada project. Tambahkan data project terlebih dahulu.")
		return
	}

	var arrKategori [NMAX]string
	var arrJumlah [NMAX]int
	var nKategori int = 0 // Menghitung jumlah kategori berbeda

	// Looping ke seluruh project yang ada
	for i := 0; i < nP; i++ {
		var ditemukan bool = false

		// Cek apakah kategori pada P[i] sudah tersimpan di arrKategori
		for j := 0; j < nKategori; j++ {
			if arrKategori[j] == P[i].kategori {
				arrJumlah[j]++   // Tambah frekuensi jumlahnya
				ditemukan = true // Tandai jika sudah ada
				break
			}
		}

		// Jika kategori belum ada, daftarkan sebagai kategori baru dengan jumlah 1
		if !ditemukan {
			arrKategori[nKategori] = P[i].kategori
			arrJumlah[nKategori] = 1
			nKategori++
		}
	}

	// Cetak statistik user
	fmt.Println("==================================")
	fmt.Println("  STATISTIK PROJECT PER KATEGORI  ")
	fmt.Println("==================================")
	for i := 0; i < nKategori; i++ {
		fmt.Printf("- %-20s : %d Project\n", arrKategori[i], arrJumlah[i])
	}
	fmt.Println("==================================")
	fmt.Println()
}
