package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_problem1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Problem 1 - Test Case pertama",
			input: `4
laptop 5 1000
phone 10 500
tablet 3 1500
laptop 3 1000`,
			expected: `laptop 3 1000 3000
tablet 3 1500 4500
laptop 5 1000 5000
phone 10 500 5000
`,
		},
		{
			name: "Problem 1 - Test Case Kedua",
			input: `3
modem 2 250
mouse 5 100
modem 1 250`,
			expected: `modem 1 250 250
modem 2 250 500
mouse 5 100 500
`,
		},
		{
			name: "Problem 1 - Test Case Input 0",
			input: `0
`,
			expected: ``,
		},
		{
			name: "Problem 1 - Test Case Single Input",
			input: `1
keyboard 10 75`,
			expected: `keyboard 10 75 750
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			inputReader := strings.NewReader(tc.input)

			outputBuffer := new(bytes.Buffer)

			problem1(inputReader, outputBuffer)

			actual := outputBuffer.String()
			if actual != tc.expected {
				t.Errorf("Test '%s' failed:\nExpected:\n%s\nGot:\n%s", tc.name, tc.expected, actual)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	sistemInventori := CreateNewInventaris()

	sistemInventori.TambahBarang("laptop", 2, 1500)
	sistemInventori.TambahBarang("mouse", 5, 100)
	sistemInventori.TambahBarang("laptop", 3, 1500)

	t.Run("Problem 2 - Test Penggabungan Stok", func(t *testing.T) {
		expectedStok := 5
		actualStok := sistemInventori.daftarBarang["laptop"].Stok
		if actualStok != expectedStok {
			t.Errorf("Ekspektasi stok laptop %d, tapi yang didapatkan %d", expectedStok, actualStok)
		}
	})

	t.Run("Problem 2 - Test Total Nilai Inventaris", func(t *testing.T) {
		expectedTotal := 8000
		actualTotal := sistemInventori.HitungTotalNilaiInventaris()
		if actualTotal != expectedTotal {
			t.Errorf("Ekspektasi total nilai %d, tapi yang didapatkan %d", expectedTotal, actualTotal)
		}
	})

	t.Run("Problem 2 - Test Sort dan Nilai Laporan", func(t *testing.T) {
		laporanAktual := sistemInventori.GetSortedLaporan()

		laporanEkspektasi := []*Barang{
			{Nama: "mouse", Stok: 5, Harga: 100, TotalNilai: 500},
			{Nama: "laptop", Stok: 5, Harga: 1500, TotalNilai: 7500},
		}

		if len(laporanAktual) != len(laporanEkspektasi) {
			t.Fatalf("Ekspektasi laporan memiliki %d item, tapi yang didapatkan %d", len(laporanEkspektasi), len(laporanAktual))
		}

		for i, barangAktual := range laporanAktual {
			barangEkspektasi := laporanEkspektasi[i]

			if barangAktual.Nama != barangEkspektasi.Nama {
				t.Errorf("Item #%d: Ekspektasi nama %s, tapi yang didapatkan %s", i+1, barangEkspektasi.Nama, barangAktual.Nama)
			}

			if barangAktual.TotalNilai != barangEkspektasi.TotalNilai {
				t.Errorf("Item #%d (%s): Ekspektasi TotalNilai %d, tapi yang didapatkan %d", i+1, barangAktual.Nama, barangEkspektasi.TotalNilai, barangAktual.TotalNilai)
			}
		}
	})
}
