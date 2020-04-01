package main

import "fmt"

func main() {
	// operator matematika
	fmt.Println(1 + 1)
	fmt.Println(5 - 1)
	fmt.Println(10 * 3)
	// pembagian
	// jika int / int, makan hasilnya akan int
	fmt.Println(10 / 4)

	// jika float/float, atau salah satu bilangan bertipe float
	// maka hasilnya akan float
	fmt.Println(10.0 / 4.0)

	// =============================================

	// Gerbang Logika
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(!true)

}
