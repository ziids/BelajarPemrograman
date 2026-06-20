package main

import ("fmt")

func main() {
	var bilangan int

	fmt.Print("Masukkan Bilangan: ")
  	fmt.Scan(&bilangan)

	if (bilangan % 2 == 0) {
		fmt.Printf("%v adalah Bilangan Genap", bilangan)
	} else {
		fmt.Printf("%v adalah Bilangan Ganjil", bilangan)
	}
}