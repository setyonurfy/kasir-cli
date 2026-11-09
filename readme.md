# Kasir CLI - Toko Kelontong

Aplikasi kasir sederhana berbasis terminal untuk mencatat pembelian barang di toko kelontong, dibuat dengan Go.

Kasir mengetik nama barang dan jumlahnya, lalu aplikasi menghitung total, memproses pembayaran, dan mencetak struk belanja yang rapi.

## Fitur

- Master data produk dengan `map[string]int` (nama barang dan harga)
- Keranjang belanja dengan slice; barang yang sama otomatis digabung
- Input kasir berulang sampai mengetik `selesai`
- Nama barang tidak peka huruf besar/kecil (`beras`, `BERAS`, dan `Beras` sama saja)
- Validasi input: barang tidak ditemukan, jumlah tidak valid, uang bayar kurang
- Uang bayar dan kembalian
- Struk dengan kolom sejajar dan format Rupiah (`Rp 1.500.000`)

## Cara menjalankan

### Tanpa Go (paling mudah)

Unduh file sesuai sistem operasi dari halaman [Releases](../../releases/latest), lalu jalankan di terminal.

| Sistem operasi | File | Cara menjalankan |
|----------------|------|------------------|
| Windows | `kasir-windows.exe` | `.\kasir-windows.exe` |
| macOS Intel | `kasir-mac-intel` | `chmod +x kasir-mac-intel && ./kasir-mac-intel` |
| macOS Apple Silicon | `kasir-mac-apple-silicon` | `chmod +x kasir-mac-apple-silicon && ./kasir-mac-apple-silicon` |
| Linux | `kasir-linux` | `chmod +x kasir-linux && ./kasir-linux` |

> File ini tidak ditandatangani, jadi Windows atau macOS bisa menampilkan peringatan keamanan. Kalau ragu, gunakan cara di bawah untuk menjalankan dari source code.

### Dengan Go (versi 1.21 atau lebih baru)

```bash
git clone https://github.com/setyonurfy/kasir-cli.git
cd kasir-cli
go run .
```

Atau pasang sebagai perintah:

```bash
go install github.com/setyonurfy/kasir-cli@latest
kasir-cli
```

## Contoh penggunaan

```
=== TOKO KELONTONG MAKMUR ===
Daftar produk:
  Beras          Rp 15.000
  Gula           Rp 14.000
  Kopi           Rp 12.000
  Minyak         Rp 20.000
  Teh             Rp 8.000
  Telur          Rp 28.000

Ketik nama barang, atau 'selesai' untuk bayar.

Nama barang : beras
Jumlah      : 2
+ Beras x2 ditambahkan.

Nama barang : minyak
Jumlah      : 1
+ Minyak x1 ditambahkan.

Nama barang : selesai

Total belanja : Rp 50.000
Uang bayar    : 100000

===========================================
           TOKO KELONTONG MAKMUR
               STRUK BELANJA
===========================================
No  Barang          Harga  Qty     Subtotal
-------------------------------------------
1   Beras       Rp 15.000    2    Rp 30.000
2   Minyak      Rp 20.000    1    Rp 20.000
-------------------------------------------
TOTAL                             Rp 50.000
BAYAR                            Rp 100.000
KEMBALIAN                         Rp 50.000
===========================================
      Terima kasih sudah berbelanja!
```

## Struktur kode

| Fungsi | Tugas |
|--------|-------|
| `main` | Alur utama: tampilkan produk, terima input, bayar, cetak struk |
| `cariProduk` | Mencari produk tanpa membedakan huruf besar/kecil |
| `tambahKeKeranjang` | Menambah barang ke keranjang, menggabungkan barang yang sama |
| `hitungTotal` | Menjumlahkan semua subtotal |
| `formatRupiah` | Mengubah angka menjadi format `Rp 1.500.000` |
| `cetakStruk` | Mencetak struk dengan kolom sejajar memakai `fmt.Printf` |

## Yang dipelajari dari project ini

- `map`, slice, dan `struct`
- Perulangan `for`, percabangan, dan fungsi dengan banyak nilai kembalian
- Membaca input terminal dengan `bufio` dan konversi dengan `strconv`
- Format specifier `fmt` (`%-10s`, `%12s`, `%d`) untuk tampilan rapi
- Cross-compile program Go untuk beberapa sistem operasi

## Rencana pengembangan

- Menyimpan riwayat transaksi ke file JSON
- Unit test untuk `formatRupiah`, `hitungTotal`, dan `tambahKeKeranjang`
- Memisahkan kode ke beberapa file

## Lisensi

Bebas dipakai untuk belajar.