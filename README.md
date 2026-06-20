# BelajarPemrograman
[![Go](https://badgen.net/badge/icon/Golang?icon=go&label)](https://github.com/ziids/BelajarPemrograman/tree/main/Go)

Soal untuk melatih Logika Dasar Dalam Pemrograman

## 1. Menghitung Total Belanja

### Kasus
Sebuah toko menjual 3 barang. Pengguna memasukkan harga dan jumlah masing-masing barang.

### Tugas
- Simpan data dalam variabel.
- Hitung total seluruh belanja.
- Tampilkan hasilnya.
  
### Contoh Input
```
Harga Barang 1 = 10000
Jumlah Barang 1 = 2

Harga Barang 2 = 5000
Jumlah Barang 2 = 3

Harga Barang 3 = 15000
Jumlah Barang 3 = 1
```

### Output
```
Total Belanja = 50000
```

## 2. Menentukan Bilangan Genap atau Ganjil
### Kasus
Pengguna memasukkan sebuah angka.

### Tugas
Tentukan apakah angka tersebut genap atau ganjil.

### Contoh Input
```
Masukkan Bilangan: 7
```

### Output
```
7 adalah Bilangan Ganjil
```

## 3. Menghitung Nilai Akhir Mahasiswa

### Kasus
Sistem akademik menghitung nilai akhir berdasarkan:
| Nilai | Persentase |
|:------|-----------:|
| Tugas | 30%        |
| UTS   | 30%        |
| UAS   | 40%        |

### Tugas
Hitung nilai akhir dan tampilkan grade:
| Nilai | Grade |
|:------|------:|
| >= 85 | A     |
| >= 70 | B     |
| >= 55 | C     |
| >= 40 | D     |
| < 40  | E     |

### Contoh Input
```
Nilai Tugas = 80
Nilai UTS = 75
Nilai UAS = 90
```

### Output
```
Nilai Akhir = 82.5
Grade = B
```

## 4. Menghitung Luas Bangun Menggunakan Function

### Kasus
Buat program yang memiliki function untuk menghitung:
- Luas Persegi
- Luas Persegi Panjang
- Luas Lingkaran

### Tugas
Gunakan function yang berbeda untuk setiap bangun.
- Gunakan loop agar menu terus muncul.
```
Program Menghitung Luas Bangun
------------------------------
1. Luas Persegi
2. Luas Persegi Panjang
3. Luas Lingkaran
0. Keluar
------------------------------
Pilih Menu: ...
```

Contoh
- `luasPersegi(5)`
- `luasPersegiPanjang(4, 6)`
- `luasLingkaran(7)`
```
Program Menghitung Luas Bangun
------------------------------
1. Luas Persegi
2. Luas Persegi Panjang
3. Luas Lingkaran
0. Keluar
------------------------------
Pilih Menu: 1

Masukkan sisi: 5
Luas Persegi: 25

Program Menghitung Luas Bangun
------------------------------
1. Luas Persegi
2. Luas Persegi Panjang
3. Luas Lingkaran
0. Keluar
------------------------------
Pilih Menu: 2

Masukkan panjang: 4
Masukkan lebar: 6
Luas Persegi Panjang: 24

Program Menghitung Luas Bangun
------------------------------
1. Luas Persegi
2. Luas Persegi Panjang
3. Luas Lingkaran
0. Keluar
------------------------------
Pilih Menu: 3

Masukkan jari-jari: 7
Luas Lingkaran: 153.86

Program Menghitung Luas Bangun
------------------------------
1. Luas Persegi
2. Luas Persegi Panjang
3. Luas Lingkaran
0. Keluar
------------------------------
Pilih Menu: 0

Berhasil Keluar Program!
```

## 5. Validasi Login Sederhana

### Kasus
Username dan password yang benar:
```
username = admin
password = 12345
```

### Tugas
- Minta input username dan password.
- Tampilkan: `"Login Berhasil"` atau `"Login Gagal"`

### Contoh Input & Output Benar
```
Masukkan Username: admin
Masukkan Password: 12345

Login Berhasil!
```

### Contoh Input & Output Salah
```
Masukkan Username: atmint
Masukkan Password: 09876

Username atau Password Salah!
```

## 6. Menghitung Faktorial

### Kasus
Faktorial suatu bilangan: `5! = 5 × 4 × 3 × 2 × 1 = 120`

### Tugas
Buat program untuk menghitung faktorial dari angka yang dimasukkan pengguna.

### Contoh Input
```
Masukkan Bilangan: 5
```

### Output
```
Faktorial: 120
```

## 7. Mencari Bilangan Terbesar dan Terkecil 3 Angka

### Kasus
Pengguna memasukkan tiga angka.

### Tugas
Tentukan angka terbesar dan terkecil.

### Contoh Input
```
Masukkan Bilangan ke-1: 12
Masukkan Bilangan ke-2: 45
Masukkan Bilangan ke-3: 30
```

### Output
```
Bilangan Terkecil yaitu bilangan ke-1: 12
Bilangan Terbesar yaitu bilangan ke-2: 45
```

## 8. Menghitung Jumlah Huruf Vokal

### Kasus
Pengguna memasukkan sebuah kalimat.

### Tugas
Hitung berapa banyak huruf vokal: `a, i, u, e, o`

### Contoh Input
```
Masukkan Kalimat: YoNdakTauKoTanyaSaya
```

### Output
```
Jumlah Vokal: 9
A: 6
I: 0
U: 1
E: 0
O: 2
```

## 9. Sistem Diskon Toko

### Kasus
Aturan diskon:
| Total Belanja | Diskon     |
|:--------------|-----------:|
| > 500000      | 20%        |
| > 250000      | 10%        |
| Lainnya       | 0%         |

### Tugas
Hitung:
- Diskon
- Total yang harus dibayar

### Contoh Input
```
Masukkan Total Belanja: 600000
```

### Output
```
Diskon = 120000 (20%)
Total Bayar = 480000
```

## 10. Sistem ATM Mini

### Kasus
Saldo awal: `0`

```
Sistem ATM Mini
--------------------
1. Cek Saldo
2. Setor Tunai
3. Tarik Tunai
0. Keluar
--------------------
Pilih Menu: ...
```

### Tugas
- Gunakan loop agar menu terus muncul.
- Gunakan function untuk setiap menu.
- Tolak penarikan jika saldo tidak cukup.

### Contoh
```
Sistem ATM Mini
--------------------
1. Cek Saldo
2. Setor Tunai
3. Tarik Tunai
0. Keluar
--------------------
Pilih Menu: 1

Sistem ATM Mini
--------------------
Saldo Anda: Rp0

Sistem ATM Mini
--------------------
1. Cek Saldo
2. Setor Tunai
3. Tarik Tunai
0. Keluar
--------------------
Pilih Menu: 2

Sistem ATM Mini
--------------------
Nominal Setor: 50000
Setor Tunai Berhasil!

Sistem ATM Mini
--------------------
1. Cek Saldo
2. Setor Tunai
3. Tarik Tunai
0. Keluar
--------------------
Pilih Menu: 3

Sistem ATM Mini
--------------------
Nominal Tarik: 1000
Tarik Tunai Berhasil!

Sistem ATM Mini
--------------------
1. Cek Saldo
2. Setor Tunai
3. Tarik Tunai
0. Keluar
--------------------
Pilih Menu: 0

Berhasil Keluar Program!
```

# Author
- Github: [@DrelezTM](https://github.com/DrelezTM) & [@ziids](https://github.com/ziids)