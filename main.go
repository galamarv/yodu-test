package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

// Kode untuk Problem 1

type Produk struct {
	Nama       string
	Jumlah     int
	Harga      int
	TotalHarga int
}

func problem1(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	//get input total produk
	n, _ := strconv.Atoi(scanner.Text())

	//create list dengan type data struct Produk dan dengan panjang sesuai jumlah produk
	semuaProduk := make([]Produk, n)

	//get input per produk
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())

		nama := parts[0]
		jumlah, _ := strconv.Atoi(parts[1])
		harga, _ := strconv.Atoi(parts[2])

		semuaProduk[i] = Produk{
			Nama:       nama,
			Jumlah:     jumlah,
			Harga:      harga,
			TotalHarga: jumlah * harga,
		}
	}

	sort.Slice(semuaProduk, func(i, j int) bool {
		// Rule 1: Sort by total penjualan dari yang paling kecil hingga terbesar.
		if semuaProduk[i].TotalHarga != semuaProduk[j].TotalHarga {
			return semuaProduk[i].TotalHarga < semuaProduk[j].TotalHarga
		}
		// Rule 2: Jika ada dua produk dengan nilai penjualan yang sama, urutkan berdasarkan nama produk secara alfabetis
		return semuaProduk[i].Nama < semuaProduk[j].Nama
	})

	// Print hasil.
	for _, item := range semuaProduk {
		fmt.Fprintf(output, "%s %d %d %d\n", item.Nama, item.Jumlah, item.Harga, item.TotalHarga)
	}
}

// kode untuk problem 2

type Barang struct {
	Nama       string
	Stok       int
	Harga      int
	TotalNilai int
}

type Inventaris struct {
	daftarBarang map[string]*Barang
}

func CreateNewInventaris() *Inventaris {
	return &Inventaris{
		daftarBarang: make(map[string]*Barang),
	}
}

func (inv *Inventaris) TambahBarang(nama string, stok int, harga int) {
	barang, eksis := inv.daftarBarang[nama]
	if eksis {
		barang.Stok += stok
		barang.TotalNilai = barang.Stok * barang.Harga
	} else {
		total := stok * harga
		inv.daftarBarang[nama] = &Barang{
			Nama:       nama,
			Stok:       stok,
			Harga:      harga,
			TotalNilai: total,
		}
	}
}

func (inv *Inventaris) HitungTotalNilaiInventaris() int {
	total := 0
	for _, barang := range inv.daftarBarang {
		total += barang.TotalNilai
	}
	return total
}

func (inv *Inventaris) GetSortedLaporan() []*Barang {
	if len(inv.daftarBarang) == 0 {
		return []*Barang{}
	}
	semuaBarang := make([]*Barang, 0, len(inv.daftarBarang))
	for _, barang := range inv.daftarBarang {
		semuaBarang = append(semuaBarang, barang)
	}
	sort.Slice(semuaBarang, func(i, j int) bool {
		return semuaBarang[i].TotalNilai < semuaBarang[j].TotalNilai
	})
	return semuaBarang
}

func (inv *Inventaris) TampilkanLaporan() {
	laporan := inv.GetSortedLaporan()
	if len(laporan) == 0 {
		fmt.Println("Inventaris kosong.")
		return
	}
	for _, barang := range laporan {
		fmt.Printf("Barang: %-10s | Stok: %-4d | Harga: %-5d | Total Nilai: %d\n",
			barang.Nama, barang.Stok, barang.Harga, barang.TotalNilai)
	}
}

func problem2Demo() {
	// Buat sistem inventaris baru.
	sistemInventaris := CreateNewInventaris()

	// Jalankan test case.
	sistemInventaris.TambahBarang("laptop", 2, 1500)
	fmt.Println("-> Menambahkan barang: laptop (2 unit, 1500)")

	sistemInventaris.TambahBarang("mouse", 5, 100)
	fmt.Println("-> Menambahkan barang: mouse (5 unit, 100)")

	sistemInventaris.TambahBarang("laptop", 3, 1500)
	fmt.Println("-> Menambahkan barang: laptop (3 unit, 1500)")

	// Hitung dan tampilkan total nilai inventaris.
	totalNilai := sistemInventaris.HitungTotalNilaiInventaris()
	fmt.Printf("\nTotal nilai keseluruhan inventaris: %d\n", totalNilai)

	// Tampilkan laporan akhir yang sudah diurutkan.
	fmt.Println("\nLaporan Inventaris")
	sistemInventaris.TampilkanLaporan()
}

func problem3() (string, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return "", fmt.Errorf("gagal membuka db: %w", err)
	}
	defer db.Close()

	createTablesSQL := `
		CREATE TABLE customers (id INT PRIMARY KEY, name VARCHAR(255));
		CREATE TABLE products (id INT PRIMARY KEY, name VARCHAR(255), category VARCHAR(255));
		CREATE TABLE transactions (id INT PRIMARY KEY, customer_id INT, product_id INT, quantity INT, price INT);
	`
	if _, err := db.Exec(createTablesSQL); err != nil {
		return "", fmt.Errorf("gagal membuat tabel: %w", err)
	}

	insertDataSQL := `
		INSERT INTO customers (id, name) VALUES (1, 'Budi'), (2, 'Sari');
		INSERT INTO products (id, name, category) VALUES (1, 'Laptop', 'Elektronik'), (2, 'Mouse', 'Elektronik'), (3, 'Buku Tulis', 'Alat Tulis');
		INSERT INTO transactions (id, customer_id, product_id, quantity, price) VALUES (1, 1, 1, 2, 10000), (2, 1, 2, 3, 500), (3, 2, 3, 10, 1000);
	`
	if _, err := db.Exec(insertDataSQL); err != nil {
		return "", fmt.Errorf("gagal memasukkan data: %w", err)
	}

	query := `
		SELECT c.name, p.category, SUM(t.quantity), SUM(t.quantity * t.price)
        FROM transactions t
        JOIN customers c ON t.customer_id = c.id
        JOIN products p ON t.product_id = p.id
        GROUP BY c.name, p.category
        ORDER BY SUM(t.quantity * t.price) DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		return "", fmt.Errorf("gagal menjalankan query: %w", err)
	}
	defer rows.Close()

	var resultBuilder strings.Builder
	for rows.Next() {
		var customerName, category string
		var totalQuantity, totalSpent int
		if err := rows.Scan(&customerName, &category, &totalQuantity, &totalSpent); err != nil {
			return "", fmt.Errorf("gagal memindai baris: %w", err)
		}
		resultBuilder.WriteString(fmt.Sprintf("%s|%s|%d|%d\n", customerName, category, totalQuantity, totalSpent))
	}

	return resultBuilder.String(), nil
}

func main() {
	//problem1(os.Stdin, os.Stdout)
	//problem2Demo()
	fmt.Print(problem3())
}
