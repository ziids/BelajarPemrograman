package main

import ("fmt")

var saldo float32 = 0

func cekSaldo() {
	fmt.Println("Sistem ATM Mini")
	fmt.Println("--------------------")
	fmt.Printf("Saldo Anda: Rp%v\n\n", saldo)
}

func setorTunai() {
	var nominal float32
	fmt.Println("Sistem ATM Mini")
	fmt.Println("--------------------")
	fmt.Print("Nominal Setor: ")
	fmt.Scan(&nominal)
	fmt.Print("\n")

	saldo += nominal
}

func tarikTunai() {
	var nominal float32
	fmt.Println("Sistem ATM Mini")
	fmt.Println("--------------------")
	fmt.Print("Nominal Tarik: ")
	fmt.Scan(&nominal)

	if (saldo < nominal) {
		fmt.Println("Tarik Tunai Gagal, Saldo Tidak Mencukupi!")
	} else {
		fmt.Println("Tarik Tunai Berhasil!")
		saldo -= nominal
	}
	fmt.Print("\n")
}

func main() {
	var menu int

	for {
		fmt.Println("Sistem ATM Mini")
		fmt.Println("--------------------")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Setor Tunai")
		fmt.Println("3. Tarik Tunai")
		fmt.Println("0. Keluar")
		fmt.Println("--------------------")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&menu)
		fmt.Print("\n")

		if (menu == 1) {
			cekSaldo()
		} else if (menu == 2) {
			setorTunai()
		} else if (menu == 3) {
			tarikTunai()
		} else if (menu == 0) {
			fmt.Println("Berhasil Keluar Program!")
			break
		} else {
			fmt.Print("Menu Tidak Valid!\n")
		}
	}
}