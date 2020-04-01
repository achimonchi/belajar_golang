package main

import "fmt"

func main() {
	fmt.Println("Belajar Golang")
	fmt.Println("Belajar" + " " + "Golang")

	// mengambil jumlah huruf
	fmt.Println(len("Belajar"))

	// mengambil character
	// namun hasilnya berupa bilangan charcode / ASCII
	fmt.Println("Belajar"[3])

	// mengambil beberapa character
	// caranya sama dgn python [index:index-akhir+1]
	// [index:] => mengambil character dari index - akhir
	fmt.Println("Belajar"[0:3])
	fmt.Println("Belajar"[3:])
}
