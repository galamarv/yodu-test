package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {
	problem1(os.Stdin, os.Stdout)
}
