package main

import (
	"fmt"
	"time"
)

func main() {
	// untuk best practice menghandle async, menggunakan channel
	done := make(chan bool)
	go hello(done)

	// lalu ambil nilai channelnya
	<-done
	fmt.Println("Ini fungsi main")
}

// function hello yang mempunyai parameter done berupa channel
func hello(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
	// lalu masukkan value kedalam channel tersebut
	done <- true
}
