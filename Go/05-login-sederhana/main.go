package main

import ("fmt")

func main() {
	var username, password string

	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)

	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	fmt.Print("\n")

	if (username == "admin" && password == "12345") {
		fmt.Println("Login Berhasil!")
	} else {
		fmt.Println("Username atau Password Salah!")
	}
}