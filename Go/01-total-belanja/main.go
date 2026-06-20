package main

import ("fmt")

func main() {
	var harga1, harga2, harga3 int
	var qty1, qty2, qty3 int

	var total int

	fmt.Print("Harga Barang 1 = ")
  	fmt.Scan(&harga1)
	fmt.Print("Jumlah Barang 1 = ")
  	fmt.Scan(&qty1)

	fmt.Print("Harga Barang 2 = ")
  	fmt.Scan(&harga2)
	fmt.Print("Jumlah Barang 2 = ")
  	fmt.Scan(&qty2)

	fmt.Print("Harga Barang 3 = ")
  	fmt.Scan(&harga3)
	fmt.Print("Jumlah Barang 3 = ")
  	fmt.Scan(&qty3)

	total = harga1 + harga2 + harga3
	fmt.Printf("Total Belanja = %v", total)
}