package main

import ("fmt")

func main() {
	var bilangan1, bilangan2, bilangan3 int

	fmt.Print("Masukkan Bilangan ke-1: ")
	fmt.Scan(&bilangan1)
	fmt.Print("Masukkan Bilangan ke-2: ")
	fmt.Scan(&bilangan2)
	fmt.Print("Masukkan Bilangan ke-3: ")
	fmt.Scan(&bilangan3)

	if (bilangan1 < bilangan2 && bilangan1 < bilangan3) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-1: %v\n", bilangan1)
	} else if (bilangan2 < bilangan1 && bilangan2 < bilangan3) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-2: %v\n", bilangan2)
	} else if (bilangan3 < bilangan1 && bilangan3 < bilangan2) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-3: %v\n", bilangan3)
	} else if (bilangan1 == bilangan2 && bilangan1 < bilangan3) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-1 & 2: %v\n", bilangan1)
	} else if (bilangan1 == bilangan3 && bilangan1 < bilangan2) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-1 & 3: %v\n", bilangan1)
	} else if (bilangan2 == bilangan3 && bilangan2 < bilangan1) {
		fmt.Printf("Bilangan Terkecil yaitu bilangan ke-2 & 3: %v\n", bilangan2)
	}

	if (bilangan1 > bilangan2 && bilangan1 > bilangan3) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-1: %v\n", bilangan1)
	} else if (bilangan2 > bilangan1 && bilangan2 > bilangan3) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-2: %v\n", bilangan2)
	} else if (bilangan3 > bilangan1 && bilangan3 > bilangan2) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-3: %v\n", bilangan3)
	} else if (bilangan1 == bilangan2 && bilangan1 > bilangan3) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-1 & 2: %v\n", bilangan1)
	} else if (bilangan1 == bilangan3 && bilangan1 > bilangan2) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-1 & 3: %v\n", bilangan1)
	} else if (bilangan2 == bilangan3 && bilangan2 > bilangan1) {
		fmt.Printf("Bilangan Terbesar yaitu bilangan ke-2 & 3: %v\n", bilangan2)
	}

	if (bilangan1 == bilangan2 && bilangan2 == bilangan3) {
		fmt.Printf("Bilangan 1, 2, 3 bernilai sama: %v\n", bilangan1)
	}
}