package main

import ("fmt")

func main() {
	var belanja, bayar, diskon float32
	var diskonPersen int

	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&belanja)

	if (belanja > 500000) {
		diskon = (belanja * 0.2)
		bayar = belanja - diskon
		diskonPersen = 20
	} else if (belanja > 250000) {
		diskon = (belanja * 0.1)
		bayar = belanja - diskon
		diskonPersen = 10
	} else {
		diskon = 0
		bayar = belanja
		diskonPersen = 0
	}

	fmt.Printf("Diskon = %v (%v%%)\n", diskon, diskonPersen)
	fmt.Printf("Total Bayar = %v", bayar)
}