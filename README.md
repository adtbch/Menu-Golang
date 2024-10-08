# Program Manajemen Pengguna Sederhana (CLI)

Ini adalah aplikasi sederhana berbasis CLI (Command Line Interface) yang ditulis menggunakan bahasa Go. Aplikasi ini menyediakan beberapa fitur seperti menampilkan "Hello, World!", melakukan operasi matematika sederhana, menyimpan dan menampilkan data pengguna, serta menghitung faktorial menggunakan fungsi rekursif. Program ini juga mencakup validasi input untuk NIM dan nomor telepon agar hanya menerima angka yang valid.

## Fitur

1. **Tampilkan "Hello, World!"**: Menampilkan pesan "Hello, World!".
2. **Operasi Matematika Sederhana**: Melakukan penjumlahan, pengurangan, perkalian, dan pembagian dua angka.
3. **Manajemen Data Pengguna**: Memungkinkan pengguna untuk memasukkan dan menampilkan data pengguna seperti nama, NIM, email, alamat, dan nomor telepon dengan validasi input.
4. **Penghitungan Faktorial**: Menghitung faktorial dari sebuah angka menggunakan fungsi rekursif.
5. **Keluar**: Menutup program.

## Cara Menggunakan

### Prasyarat
- Pastikan Anda telah menginstall [Go](https://golang.org/dl/) pada sistem Anda.
- Clone repository ini ke komputer lokal Anda:
  ```bash
  git clone https://github.com/adtbch/Menu-Golang
  ```

### Menjalankan Program

1. Masuk ke direktori project:
   ```bash
   cd nama-repo
   ```
2. Jalankan program dengan perintah berikut:
   ```bash
   go run main.go
   ```

### Opsi Menu Utama

Setelah menjalankan program, Anda akan disajikan dengan menu. Anda dapat memilih opsi berikut dengan memasukkan angka sesuai:

1. **Tampilkan "Hello, World!"**: Menampilkan "Hello, World!" pada layar.
2. **Operasi Matematika Sederhana**: Meminta pengguna untuk memasukkan dua angka, kemudian melakukan operasi penjumlahan, pengurangan, perkalian, dan pembagian.
3. **Manajemen Data Pengguna**: Memungkinkan Anda untuk memasukkan data pengguna, termasuk nama, NIM (dengan validasi angka), email (dengan validasi format), alamat, dan nomor telepon (dengan validasi angka). Data tersebut kemudian akan ditampilkan.
4. **Penghitungan Faktorial**: Meminta pengguna untuk memasukkan angka, kemudian menghitung dan menampilkan faktorial dari angka tersebut menggunakan fungsi rekursif.
5. **Keluar**: Mengakhiri program.

### Contoh Input Data Pengguna

Saat memilih opsi **Manajemen Data Pengguna**, Anda akan diminta untuk memasukkan:
- **Nama**: Input teks bebas.
- **NIM**: Harus berupa angka.
- **Email**: Harus menggunakan format email yang valid (mengandung '@' dan '.').
- **Alamat**: Input teks bebas.
- **Nomor HP**: Harus berupa angka.

Jika Anda mencoba memasukkan NIM, email, atau nomor telepon yang tidak valid, program akan menampilkan pesan kesalahan yang sesuai.

### Struktur Program

- **Menu Utama**: Sebuah loop yang menampilkan menu untuk navigasi ke berbagai fitur program.
- **Operasi Matematika**: Fungsi untuk melakukan operasi dasar matematika (penjumlahan, pengurangan, perkalian, dan pembagian).
- **Manajemen Data Pengguna**: Menggunakan slice dari map untuk menyimpan beberapa data pengguna, setiap pengguna berisi nama, NIM, email, alamat, dan nomor telepon.
- **Fungsi Faktorial**: Menghitung faktorial dari sebuah angka menggunakan rekursi.
- **Validasi Input**: Memastikan bahwa NIM dan nomor telepon hanya berisi angka, dan email mengikuti format yang benar.

### Struktur Kode

Program ini terdiri dari beberapa fungsi:
- `tampilkanHelloWorld()`: Menampilkan pesan "Hello, World!".
- `operasiMatematika()`: Melakukan operasi matematika sederhana.
- `kelolaDataPengguna()`: Mengelola input data pengguna, memvalidasi NIM dan nomor telepon, serta menampilkan data pengguna.
- `tampilkanDataPengguna()`: Menampilkan semua data pengguna yang telah tersimpan.
- `hitungFaktorial()`: Meminta angka dari pengguna dan menghitung faktorialnya.
- `faktorial(n int)`: Menghitung faktorial dari sebuah angka menggunakan fungsi rekursif.


## Kontak

Untuk pertanyaan atau masalah terkait proyek ini, Anda dapat menghubungi saya di: adit.bachtiar091@gmail.com

---
