package main

import "fmt"

func main() {
	// String
	fmt.Println("Ini String")

	// integer
	// uint : tidak boleh ada negatif
	// int : boleh negatif
	// int8, int16, int32, int 64 => ukuran kapasistas integer
	// byte => alias dari uint8
	// int => alias dari int32 atau int64
	fmt.Print("Contoh integer : ")
	fmt.Println(1)

	// float
	// float32, float64 => ukuran nilai
	// menggunakan . (point)
	fmt.Print("Contoh float : ")
	fmt.Println(1.1)

	// booleans
	// tipe data yg berisi TRUE atau FALSE
	fmt.Print("Contoh Boolean : ")
	fmt.Println(true)
}
