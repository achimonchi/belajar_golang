package main

import "fmt"

func main() {

	flag := true

	// untuk if else, sedikit ada perbedaan.
	// jika biasanya if(kondisi) {}, maka di golang tidak menggunakan ()
	// melainkan if kondisi {}
	if flag == false {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	// dan pada else, itu harus sebaris dengan } pada if
	// } else {}
	// tidak bisa seperti ini...
	// if flag == false {
	// 	fmt.Println(false)
	// }
	// else {
	// 	fmt.Println(true)
	// }

	// kombinasi antara if dan for
	fmt.Println("")
	fmt.Println("=========================")
	i := 1
	for i <= 30 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FooBar")
		} else if i%3 == 0 {
			fmt.Println("Foo")
		} else if i%5 == 0 {
			fmt.Println("Bar")
		} else {
			fmt.Println(i)
		}
		i++
	}
}
