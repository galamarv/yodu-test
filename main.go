package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// kode untuk problem 3

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

// kode untuk problem 4

type ElasticsearchDoc struct {
	ProductName  string `bson:"product_name" json:"product_name"`
	CustomerName string `bson:"customer_name" json:"customer_name"`
	Category     string `bson:"category" json:"category"`
	TotalQty     int    `bson:"total_qty" json:"total_qty"`
	TotalRevenue int    `bson:"total_revenue" json:"total_revenue"`
}

func problem4(mongoURI string) (string, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return "", fmt.Errorf("gagal konek ke MongoDB: %w", err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("toko")
	salesColl := db.Collection("sales")
	productsColl := db.Collection("products")
	customersColl := db.Collection("customers")

	salesColl.DeleteMany(context.TODO(), bson.M{})
	productsColl.DeleteMany(context.TODO(), bson.M{})
	customersColl.DeleteMany(context.TODO(), bson.M{})

	p1, _ := primitive.ObjectIDFromHex("60c72b2f5f1b2c6a8c7a6b1a")
	p2, _ := primitive.ObjectIDFromHex("60c72b2f5f1b2c6a8c7a6b1b")
	p3, _ := primitive.ObjectIDFromHex("60c72b2f5f1b2c6a8c7a6b1c")
	c1, _ := primitive.ObjectIDFromHex("60c72b3a5f1b2c6a8c7a6b1d")
	c2, _ := primitive.ObjectIDFromHex("60c72b3a5f1b2c6a8c7a6b1e")

	productsData := []interface{}{
		bson.M{"_id": p1, "name": "mouse", "category": "Elektronik"},
		bson.M{"_id": p2, "name": "keyboard", "category": "Elektronik"},
		bson.M{"_id": p3, "name": "monitor", "category": "Elektronik"},
	}
	customersData := []interface{}{
		bson.M{"_id": c1, "name": "Budi"},
		bson.M{"_id": c2, "name": "Sari"},
	}
	salesData := []interface{}{
		bson.M{"product_id": p1, "customer_id": c1, "quantity": 60, "unit_price": 100, "created_at": time.Now()},
		bson.M{"product_id": p1, "customer_id": c1, "quantity": 50, "unit_price": 100, "created_at": time.Now()},
		bson.M{"product_id": p2, "customer_id": c1, "quantity": 120, "unit_price": 500, "created_at": time.Now()},
		bson.M{"product_id": p3, "customer_id": c2, "quantity": 90, "unit_price": 800, "created_at": time.Now()},
	}

	productsColl.InsertMany(context.TODO(), productsData)
	customersColl.InsertMany(context.TODO(), customersData)
	salesColl.InsertMany(context.TODO(), salesData)

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.M{"from": "products", "localField": "product_id", "foreignField": "_id", "as": "product_info"}}},
		{{"$unwind", "$product_info"}},
		{{"$lookup", bson.M{"from": "customers", "localField": "customer_id", "foreignField": "_id", "as": "customer_info"}}},
		{{"$unwind", "$customer_info"}},
		{{"$group", bson.M{
			"_id":           bson.M{"product_id": "$product_info._id", "customer_id": "$customer_info._id"},
			"product_name":  bson.M{"$first": "$product_info.name"},
			"customer_name": bson.M{"$first": "$customer_info.name"},
			"category":      bson.M{"$first": "$product_info.category"},
			"total_qty":     bson.M{"$sum": "$quantity"},
			"total_revenue": bson.M{"$sum": bson.M{"$multiply": bson.A{"$quantity", "$unit_price"}}},
		}}},
		{{"$match", bson.M{"total_qty": bson.M{"$gt": 100}}}},
		{{"$sort", bson.M{"total_revenue": -1}}},
		{{"$project", bson.M{
			"_id": 0, "product_name": "$product_name", "customer_name": "$customer_name",
			"category": "$category", "total_qty": "$total_qty", "total_revenue": "$total_revenue",
		}}},
	}

	cursor, err := salesColl.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return "", fmt.Errorf("agregasi gagal: %w", err)
	}
	defer cursor.Close(context.TODO())

	var results []ElasticsearchDoc
	if err = cursor.All(context.TODO(), &results); err != nil {
		return "", fmt.Errorf("gagal membaca hasil: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", fmt.Errorf("gagal marshaling JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func main() {
	//problem1(os.Stdin, os.Stdout)
	//problem2Demo()
	fmt.Print(problem3())
}
