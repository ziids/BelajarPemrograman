package main

import ("fmt")

func luasPersegi(sisi float32) float32 {
	return sisi * sisi
}

func luasPersegiPanjang(panjang float32, lebar float32) float32 {
	return panjang * lebar
}

func luasLingkaran(jariJari float32) float32 {
	return 3.14 * jariJari * jariJari
}

func main() {
	var menu int

	for {
		fmt.Println("Program Menghitung Luas Bangun")
		fmt.Println("------------------------------")
		fmt.Println("1. Luas Persegi")
		fmt.Println("2. Luas Persegi Panjang")
		fmt.Println("3. Luas Lingkaran")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------------")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&menu)
		fmt.Print("\n")

		if (menu == 1) {
			var sisi float32
			fmt.Print("Masukkan sisi: ")
			fmt.Scan(&sisi)

			luas := luasPersegi(sisi)
			fmt.Printf("Luas Persegi: %.2f\n\n", luas)
		} else if (menu == 2) {
			var panjang, lebar float32
			fmt.Print("Masukkan panjang: ")
			fmt.Scan(&panjang)
			fmt.Print("Masukkan lebar: ")
			fmt.Scan(&lebar)

			luas := luasPersegiPanjang(panjang, lebar)
			fmt.Printf("Luas Persegi Panjang: %.2f\n\n", luas)
		} else if (menu == 3) {
			var jariJari float32
			fmt.Print("Masukkan jari-jari: ")
			fmt.Scan(&jariJari)

			luas := luasLingkaran(jariJari)
			fmt.Printf("Luas Lingkaran: %.2f\n\n", luas)
		} else if (menu == 0) {
			fmt.Println("Berhasil Keluar Program!")
			break
		} else {
			fmt.Println("Menu Tidak Valid!\n")
		}
	}
}