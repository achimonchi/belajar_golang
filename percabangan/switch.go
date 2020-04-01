package main

import (
	"fmt"
	"runtime"
)

func main() {
	flag := true

	switch flag {
	case true:
		fmt.Println("Yap Benar !")
	case false:
		fmt.Println("Sayang Sekali T_T")
	}

	i := 1
	for i <= 10 {
		switch i {
		case 1:
			fmt.Println("Satu")
		case 2:
			fmt.Println("Dua")
		case 3:
			fmt.Println("Tiga")
		default:
			fmt.Println("Maaf, kepalaku sudah pusing !")
		}
		i++
	}

	sistemOperasi := runtime.GOOS

	switch sistemOperasi {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	}
}
