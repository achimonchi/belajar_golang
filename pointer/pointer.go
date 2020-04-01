package main

import "fmt"

func changePoint(point *int) {
	*point = 200
}

func main() {
	// pointer
	// adalah untuk mengambil alamat memory
	// * artinya untuk mengambil nilai dari pointer yang ditunjukkan si pointer
	// & artinya untuk mengambil alamat memory
	var point = 100
	var pointAddr *int = &point

	fmt.Println(point)
	fmt.Println(pointAddr)
	fmt.Println("Nilai dari pointAddr adalah", *pointAddr)

	// jika mengubah nilai berdasarkan pointernya, maka
	// variable yg dituju si pointer juga akan berubah
	fmt.Println()
	*pointAddr = 300
	fmt.Println(point)
	fmt.Println(pointAddr)
	fmt.Println("Nilai dari pointAddr adalah", *pointAddr)

	// untuk merubahnya, dapat menggunakan function seperti ini
	// value point akan berubah, karena pada function changePoint
	// ia merubah nilai point berdasarkan alamat memori
	// jadi, walau variable point bukanlah global variable,
	// maka tetap bisa diubah menggunakan pointer
	fmt.Println()
	changePoint(&point)
	fmt.Println(point)
	fmt.Println(pointAddr)
	fmt.Println("Nilai dari pointAddr adalah", *pointAddr)
}
