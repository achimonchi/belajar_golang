package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine adalah suatu cara untuk mengeksekusi
	// dengan asynchronus
	go HelloGo()

	// jika tidak membuat sleep / menunggu selama 1 detik
	// maka helloGo tidak akan nampil
	// namun, hal ini bukanlah suatu best practice
	time.Sleep(1 * time.Second)

	fmt.Println("Ini main function")
}

func HelloGo() {
	fmt.Println("Hello from Go")
}
