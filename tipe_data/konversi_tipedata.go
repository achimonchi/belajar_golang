package main

import (
	"fmt"
	"strconv"
)

func main() {
	// konversi int -> int
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	// konversi int -> float
	var z float32 = float32(y) / 4
	fmt.Println(z)

	// konversi float -> int
	// jika float nya berkoma, maka tidak akan dibulatkan
	// ia hanya mengambil nilai awalnya, dan koma akan diabaikan
	var a int = int(z)
	fmt.Println(a)

	// untuk konversi dari string ke int ataupun sebaliknya,
	// menggunakan library bawaan dari golang, yaitu strconv
	// Atoi => string to int
	// Itoa => int to string
	// dan return dari method ini ada 2 buah, yaitu
	// hasil conversi dan error

	var nilai string = "100"
	// var nilaiInt int, error = strconv.Atoi(nilai)

	// jika dengan menggunakan cara diatas, maka kita mendeklarasikan variable error
	// padahal,kita tidak membutuhkannya. maka bisa dilakukan seperti ini
	nilaiInt, _ := strconv.Atoi(nilai)
	fmt.Println(nilaiInt + 30)

}
