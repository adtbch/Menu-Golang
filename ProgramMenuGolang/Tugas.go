package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Slice untuk menyimpan data pengguna
var dataPengguna []map[string]string

func main() {
	var menu int

	// Loop untuk menampilkan menu utama
	for {
		fmt.Println("\n=== Menu Utama ===")
		fmt.Println("1. Tampilkan Hello World")
		fmt.Println("2. Operasi Matematika Sederhana")
		fmt.Println("3. Simpan dan Tampilkan Data Pengguna")
		fmt.Println("4. Hitung Faktorial (Recursive Function)")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu (1-5): ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			tampilkanHelloWorld()
		case 2:
			operasiMatematika()
		case 3:
			kelolaDataPengguna()
		case 4:
			hitungFaktorial()
		case 5:
			fmt.Println("Terima kasih! Program akan ditutup.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Fungsi untuk menampilkan "Hello World!"
func tampilkanHelloWorld() {
	fmt.Println("\nHello, World!")
}

// Fungsi untuk operasi matematika sederhana
func operasiMatematika() {
	var angkaPertama, angkaKedua float64
	fmt.Print("\nMasukkan angka pertama: ")
	fmt.Scan(&angkaPertama)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&angkaKedua)

	penjumlahan := angkaPertama + angkaKedua
	pengurangan := angkaPertama - angkaKedua
	perkalian := angkaPertama * angkaKedua
	pembagian := 0.0
	if angkaKedua != 0 {
		pembagian = angkaPertama / angkaKedua
	} else {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
	}

	// Tampilkan hasil
	fmt.Printf("\nHasil Penjumlahan: %.2f\n", penjumlahan)
	fmt.Printf("Hasil Pengurangan: %.2f\n", pengurangan)
	fmt.Printf("Hasil Perkalian: %.2f\n", perkalian)
	if angkaKedua != 0 {
		fmt.Printf("Hasil Pembagian: %.2f\n", pembagian)
	}
}

// Fungsi untuk menyimpan dan menampilkan data pengguna
func kelolaDataPengguna() {
	data := make(map[string]string)
	var input string

	// Input data pengguna
	fmt.Print("\nMasukkan Nama (atau ketik 'selesai' untuk berhenti): ")
	fmt.Scan(&input)
	if strings.ToLower(input) == "selesai" {
		// Jangan simpan data jika pengguna langsung mengetik 'selesai' di Nama
		fmt.Println("Tidak ada data yang disimpan.")
		return
	}
	data["Nama"] = input

	for {
		fmt.Print("Masukkan NIM (atau ketik 'selesai' untuk berhenti): ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "selesai" {
			// Simpan data yang sudah diinput (sampai Nama)
			dataPengguna = append(dataPengguna, data)
			tampilkanDataPengguna()
			return
		}

		// Anonymous function untuk memvalidasi NIM (hanya angka)
		isValidNIM := func(nim string) bool {
			_, err := strconv.Atoi(nim) // validasi apakah input hanya angka
			return err == nil
		}

		// Validasi NIM dengan anonymous function
		if !isValidNIM(input) {
			fmt.Println("NIM tidak valid. Harus berupa angka. Data tidak disimpan.")
		} else {
			data["NIM"] = input
			break
		}
	}

	for {
		fmt.Print("Masukkan Email (atau ketik 'selesai' untuk berhenti): ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "selesai" {
			// Simpan data yang sudah diinput (sampai NIM)
			dataPengguna = append(dataPengguna, data)
			tampilkanDataPengguna()
			return
			// break
		}
		// Anonymous function untuk memvalidasi email
		isValidEmail := func(email string) bool {
			// Sederhana: email dianggap valid jika mengandung '@' dan '.'
			return strings.Contains(email, "@") && strings.Contains(email, ".")
		}

		// Validasi email dengan anonymous function
		if !isValidEmail(input) {
			fmt.Println("Email tidak valid. Data tidak disimpan.")
		} else {
			data["Email"] = input
			break
		}
	}

	fmt.Print("Masukkan Alamat (atau ketik 'selesai' untuk berhenti): ")
	fmt.Scan(&input)
	if strings.ToLower(input) == "selesai" {
		// Simpan data yang sudah diinput (sampai Email)
		dataPengguna = append(dataPengguna, data)
		tampilkanDataPengguna()
		return
	}
	data["Alamat"] = input

	for {
		fmt.Print("Masukkan Nomor HP (atau ketik 'selesai' untuk berhenti): ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "selesai" {
			// Simpan data yang sudah diinput (sampai Alamat)
			dataPengguna = append(dataPengguna, data)
			tampilkanDataPengguna()
			return
		}

		// Anonymous function untuk memvalidasi NoHP (hanya angka)
		isValidNoHP := func(noHP string) bool {
			_, err := strconv.Atoi(noHP) // validasi apakah input hanya angka
			return err == nil
		}

		// Validasi nomor HP dengan anonymous function
		if !isValidNoHP(input) {
			fmt.Println("Nomor HP tidak valid. Harus berupa angka. Data tidak disimpan.")
		} else {
			data["NoHP"] = input
			break
		}
	}

	// Simpan data pengguna lengkap ke dalam slice
	dataPengguna = append(dataPengguna, data)
	fmt.Println("Data berhasil disimpan.")
	tampilkanDataPengguna()
}

// Fungsi untuk menampilkan semua data pengguna
func tampilkanDataPengguna() {
	if len(dataPengguna) == 0 {
		fmt.Println("\nBelum ada data pengguna.")
		return
	}

	fmt.Println("\n=== Data Pengguna ===")
	for i, pengguna := range dataPengguna {
		fmt.Printf("Pengguna %d:\n", i+1)
		for key, value := range pengguna {
			fmt.Printf("%s: %s\n", key, value)
		}
		fmt.Println()
	}
}

// Fungsi rekursif untuk menghitung faktorial
func hitungFaktorial() {
	var angka int
	fmt.Print("Masukkan angka untuk menghitung faktorial: ")
	fmt.Scan(&angka)

	result := faktorial(angka)
	fmt.Printf("Faktorial dari %d adalah %d\n", angka, result)
}

// Fungsi faktorial rekursif
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}
