package main

import "fmt"

// global variabel
var global = "Ini Variable Golbal"

func main() {

	var hello = "Hello World"
	fmt.Println(hello)

	hello = "Reyhan"
	fmt.Println(hello)

	var nama string
	nama = "Jovie"
	fmt.Println(nama)

	// pemanggilan ini hanya bisa di lakukan di dalam function / method

	namaLengkap := "Reyhan Jovie"
	fmt.Println(namaLengkap)

	// int
	var nilai int8
	nilai = 100
	fmt.Println(nilai)

	// jika tidak di definisikan, maka
	// akan bernilai int (cek di file tipe data)
	nilai2 := 120
	fmt.Println(nilai2)

	fmt.Println(global)
	x()

	// variable constan
	const company string = "NooBee"
	fmt.Println(company)

	// setelah di definisikan, constant variable tidak bisa
	// di ubah lagi isinya
	// company = "New NooBee"
	// fmt.Println(company)
}

func x() {
	fmt.Println(global)
}
