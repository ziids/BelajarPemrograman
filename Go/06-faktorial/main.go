package main

import ("fmt")

func main() {
	var bilangan int
	var faktorial int = 1

	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&bilangan)

	for i := bilangan; i > 0; i-- {
		faktorial *= i
	}

	fmt.Printf("Faktorial: %v", faktorial)
}