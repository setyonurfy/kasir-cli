package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	lebar    = 43 // lebar struk dalam karakter
	namaToko = "TOKO KELONTONG MAKMUR"
)

// Item adalah satu baris di keranjang belanja
type Item struct {
	Nama     string
	Harga    int
	Jumlah   int
	Subtotal int
}

func main() {
	// Master data produk: nama -> harga
	produk := map[string]int{
		"Beras":  15000,
		"Minyak": 20000,
		"Gula":   14000,
		"Telur":  28000,
		"Kopi":   12000,
		"Teh":    8000,
	}

	// Slice untuk keranjang belanja
	var keranjang []Item
	reader := bufio.NewReader(os.Stdin)

	tampilkanProduk(produk)
	fmt.Println("\nKetik nama barang, atau 'selesai' untuk bayar.")

	// Perulangan input barang sampai kasir mengetik "selesai"
	for {
		fmt.Print("\nNama barang : ")
		input, ok := bacaBaris(reader)
		if !ok || strings.EqualFold(input, "selesai") {
			break
		}
		if input == "" {
			continue
		}

		nama, harga, ditemukan := cariProduk(produk, input)
		if !ditemukan {
			fmt.Printf("Barang %q tidak ditemukan.\n", input)
			continue
		}

		fmt.Print("Jumlah      : ")
		jmlStr, ok := bacaBaris(reader)
		if !ok {
			break
		}
		jumlah, err := strconv.Atoi(jmlStr)
		if err != nil || jumlah <= 0 {
			fmt.Println("Jumlah harus berupa angka lebih dari 0.")
			continue
		}

		keranjang = tambahKeKeranjang(keranjang, nama, harga, jumlah)
		fmt.Printf("+ %s x%d ditambahkan.\n", nama, jumlah)
	}

	if len(keranjang) == 0 {
		fmt.Println("\nKeranjang kosong, tidak ada struk.")
		return
	}

	total := hitungTotal(keranjang)

	// Proses pembayaran: ulangi sampai uang cukup
	var bayar int
	for {
		fmt.Printf("\nTotal belanja : %s\n", formatRupiah(total))
		fmt.Print("Uang bayar    : ")
		bayarStr, ok := bacaBaris(reader)
		if !ok {
			return
		}
		n, err := strconv.Atoi(bayarStr)
		if err != nil || n <= 0 {
			fmt.Println("Masukkan nominal berupa angka.")
			continue
		}
		if n < total {
			fmt.Printf("Uang kurang %s.\n", formatRupiah(total-n))
			continue
		}
		bayar = n
		break
	}

	cetakStruk(keranjang, total, bayar)
}
func bacaBaris(r *bufio.Reader) (string, bool) {
	s, err := r.ReadString('\n')
	if err != nil && s == "" {
		return "", false
	}
	return strings.TrimSpace(s), true
}

func cariProduk(produk map[string]int, input string) (string, int, bool) {
	for nama, harga := range produk {
		if strings.EqualFold(nama, input) {
			return nama, harga, true
		}
	}
	return "", 0, false
}

func tampilkanProduk(produk map[string]int) {
	fmt.Printf("=== %s ===\n", namaToko)
	fmt.Println("Daftar produk:")

	nama := make([]string, 0, len(produk))
	for n := range produk {
		nama = append(nama, n)
	}
	sort.Strings(nama)

	for _, n := range nama {
		fmt.Printf(" %-10s %12s\n", n, formatRupiah(produk[n]))
	}
}

func tambahKeKeranjang(keranjang []Item, nama string, harga, jumlah int) []Item {
	for i := range keranjang {
		if keranjang[i].Nama == nama {
			keranjang[i].Jumlah += jumlah
			keranjang[i].Subtotal = keranjang[i].Harga * keranjang[i].Jumlah
		}
	}
	return append(keranjang, Item{
		Nama:     nama,
		Harga:    harga,
		Jumlah:   jumlah,
		Subtotal: harga * jumlah,
	})
}

func hitungTotal(keranjang []Item) int {
	total := 0
	for _, item := range keranjang {
		total += item.Subtotal
	}
	return total
}

func formatRupiah(n int) string {
	s := strconv.Itoa(n)
	var sb strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			sb.WriteByte('.')
		}
		sb.WriteRune(c)
	}
	return "Rp " + sb.String()
}

func tengah(teks string) string {
	pad := (lebar - len(teks)) / 2
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat(" ", pad) + teks
}

func cetakStruk(keranjang []Item, total, bayar int) {
	sama := strings.Repeat("=", lebar)
	putus := strings.Repeat("-", lebar)

	fmt.Println()
	fmt.Println(sama)
	fmt.Println(tengah(namaToko))
	fmt.Println(tengah("STRUK BELANJA"))
	fmt.Println(sama)
	fmt.Printf("%-3s %-10s %10s %4s %12s\n", "No", "Barang", "Harga", "Qty", "Sutotal")
	fmt.Println(putus)

	for i, item := range keranjang {
		fmt.Printf("%-3d %-10s %10s %4d %12s\n", i+1, item.Nama, formatRupiah(item.Harga), item.Jumlah, formatRupiah(item.Subtotal))
	}
	fmt.Println(putus)
	fmt.Printf("%-3s %12s\n", "TOTAL", formatRupiah(total))
	fmt.Printf("%-3s %12s\n", "TOTAL", formatRupiah(bayar))
	fmt.Printf("%-30s %12s\n", "KEMBALIAN", formatRupiah(bayar-total))
	fmt.Println(sama)
	fmt.Println(tengah("Terima kasih sudah berbelanja"))
}
