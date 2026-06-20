package main

import (
	"fmt"
	"strings"
)

func main() {
	var kalimat string
	var a, i, u, e, o, jumlah int

	fmt.Print("Masukkan Kata: ")
	fmt.Scan(&kalimat)

	splitter := strings.Split(kalimat, "")

	for index := 0; index < len(splitter); index++ {
		if (splitter[index] == "a" || splitter[index] == "A") {
			a++
			jumlah++
		} else if (splitter[index] == "i" || splitter[index] == "I") {
			i++
			jumlah++
		} else if (splitter[index] == "u" || splitter[index] == "U") {
			u++
			jumlah++
		} else if (splitter[index] == "e" || splitter[index] == "E") {
			e++
			jumlah++
		} else if (splitter[index] == "o" || splitter[index] == "O") {
			o++
			jumlah++
		}
	}

	fmt.Printf("Jumlah Vokal: %v\n", jumlah)
	fmt.Printf("A: %v\n", a)
	fmt.Printf("I: %v\n", i)
	fmt.Printf("U: %v\n", u)
	fmt.Printf("E: %v\n", e)
	fmt.Printf("O: %v\n", o)
}