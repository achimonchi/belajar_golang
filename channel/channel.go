package main

import "fmt"

// untuk menghandle proses async, dibutuhkan channel
// pada function yg dilakukan secara async, akan dikirimkan
// channel seperti dibawah ini
// done <- true (sender)
// lalu, jika selesai, untuk melanjutkan proses dilakukan receiver seperti ini
// <- done (receiver)
func main() {
	// ini adalah channel buffer
	hello := make(chan string, 2)

	// mengisi nilai channel buffer
	hello <- "Hello"
	hello <- "Golang"
	close(hello)
	// untuk menampilkannya, dapat dilakukan seperti ini
	fmt.Println("manual", <-hello)
	// fmt.Println("manual", <-hello)

	// jika menampilkannya ingin menggunakan looping, maka harus di close dulu
	// channel buffer nya.
	// dan jika isi buffernya ada 2, lalu telah di tampilkan secara manual ke2 nya
	// maka pada looping, tidak ada yang ditampilkan.
	for v := range hello {
		fmt.Println(v)
	}
}
