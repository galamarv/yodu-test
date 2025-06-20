# Solusi Go untuk Tes Yodu

Repositori ini berisi solusi lengkap dalam bahasa Go 

* **Problem 1**: Inventory Value Sorting.
* **Problem 2**: OOP Inventory System.
* **Problem 3**: Business Report with JOIN (MySQL).
* **Problem 4**: MongoDB Aggregation with Join for Elasticsearch.

## Struktur Proyek

* `main.go`: Berisi kode sumber untuk semua solusi (Problem 1, 2, 3, dan 4).
* `main_test.go`: Berisi unit test yang komprehensif untuk memvalidasi keempat solusi.
* `go.mod` & `go.sum`: File modul Go yang mengelola dependensi proyek.

## Panduan Penggunaan

### 1. Persiapan Awal

**Kloning Repositori**

Pertama, kloning repositori ini ke komputer lokal Anda menggunakan Git.

```sh
git clone <URL_REPOSITORI>
cd <NAMA_DIREKTORI_PROYEK>
```

**Prasyarat**

* Pastikan Anda sudah menginstal **Go** (versi 1.22.5 atau lebih baru).

**Instalasi Dependensi**

Proyek ini memerlukan beberapa pustaka eksternal (untuk SQLite dan MongoDB). Jalankan perintah berikut untuk mengunduh dan menginstalnya secara otomatis.

```sh
go mod tidy
```

### 2. Menjalankan Unit Test (Cara Verifikasi Utama)

Ini adalah cara terbaik dan tercepat untuk memverifikasi bahwa semua solusi (Problem 1, 2, 3, dan 4) berfungsi dengan benar.

Dari direktori utama proyek, jalankan perintah:

```sh
go test -v
```

**Hasil yang Diharapkan**

Jika semua solusi benar, Anda akan melihat output yang menunjukkan bahwa semua test case **PASS**.

```
=== RUN   Test_problem1
=== RUN   Test_problem1/Problem_1_-_Test_Case_pertama
--- PASS: Test_problem1 (0.00s)
    --- PASS: Test_problem1/Problem_1_-_Test_Case_pertama (0.00s)
    ...
=== RUN   TestSistemInventaris
=== RUN   TestSistemInventaris/Problem_2_-_Test_Penggabungan_Stok
--- PASS: TestSistemInventaris (0.00s)
    --- PASS: TestSistemInventaris/Problem_2_-_Test_Penggabungan_Stok (0.00s)
    ...
=== RUN   Test_problem3
--- PASS: Test_problem3 (0.00s)
=== RUN   Test_problem4
=== RUN   Test_problem4
--- PASS: Test_problem4 (0.01s)
    --- PASS: Test_problem4 (0.00s)
PASS
ok      yodu-test    0.350s
```

### 3. Menjalankan Demo per Soal

Anda dapat menjalankan demonstrasi untuk masing-masing soal (1, 2, atau 3) dengan mengedit fungsi `main` di dalam file `main.go`.

#### Menjalankan Demo Problem 1 (Input Manual)

1. Buka file `main.go`.
2. Ubah fungsi `main` menjadi seperti ini:
   ```go
   func main() {
       problem1(os.Stdin, os.Stdout)
   }
   ```
3. Simpan file, lalu jalankan program dari terminal:
   ```sh
   go run main.go
   ```
4. Program akan menunggu. Ketik input berikut ke terminal:
   ```
   4
   laptop 5 1000
   phone 10 500
   tablet 3 1500
   laptop 3 1000
   ```
5. Program akan mencetak output yang sudah diurutkan.

#### Menjalankan Demo Problem 2 (Inventaris OOP)

1. Buka file `main.go`.
2. Ubah fungsi `main` menjadi seperti ini:
   ```go
   func main() {
       problem2Demo()
   }
   ```
3. Simpan file, lalu jalankan program dari terminal:
   ```sh
   go run main.go
   ```
4. Program akan langsung mencetak hasil demonstrasi sistem inventaris.

#### Menjalankan Demo Problem 3 (SQL Query)

1. Buka file `main.go`.
2. Ubah fungsi `main` menjadi seperti ini:
   ```go
   func main() {
       fmt.Print(problem3())
   }
   ```
3. Simpan file, lalu jalankan program dari terminal:
   ```sh
   go run main.go
   ```
4. Program akan menjalankan query pada database SQLite in-memory dan mencetak hasilnya.
