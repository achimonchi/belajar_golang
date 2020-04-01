package main

import "fmt"

func main() {
	// jika melakukan hal ini, maka
	// x akan memiliki jumlah array
	// yang tetap yaitu 5 data
	var x [5]int

	x[0] = 100
	x[1] = 130
	x[2] = 38
	x[3] = 95
	x[4] = 300

	// jika dibuat x[5], maka akan error array index out of bond
	// yang mana artinya, arraynya melebih batas
	// x[5] = 10

	fmt.Println(x)

	// secara garis besar, array di Golang sama dengan array di java
	// php, js, atau yang lainnya. Yaitu index dimulai dari 0

	// untuk menampilkan semua isi array 1 per 1, maka bisa dilakukan perulangan
	// seperti biasa. yaitu menggunakan len() untuk mengambil jumlah elemen di array
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// adapun bisa digunakan juga foreach versi golang
	// index bisa di ganti dengan variable lain.
	// range x adalah array
	// for index, value := range x {
	// 	fmt.Println("index ", index, " = ", value)
	// }

	// namun, jika tidak membutuhkan index, bisa menggunakan
	// _ (underscore).
	for _, value := range x {
		fmt.Println("isinya adalah : ", value)
	}

	// lalu ada cara membuat array yang lebih simple,
	// yaitu :

	names := [5]string{
		"Reyhan",
		"Jovie",
		"Dwiputra",
		"NooBee",
		"Achimonchi",
	}

	fmt.Println(names)
}
