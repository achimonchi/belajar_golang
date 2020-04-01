package main

import "fmt"

func main() {
	i := 1
	// for biasa
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// fmt.Println("Looping infinite")

	// ini mirip seperti while(true)
	// akan membuat looping for infinite
	// for {
	// 	fmt.Println("Looping infinite >_<")
	// }
	fmt.Println("======================================")
	fmt.Println("Looping infinite with break point")
	fmt.Println("")
	// untuk melakukan break pada looping infinite
	flag := false
	for {
		if flag == true {
			break
		}

		if i == 10 {
			flag = true
		}

		fmt.Println("Looping with break ", i)

		i++
	}

	// dapat juga melakukan looping for biasa
	fmt.Println("======================================")
	fmt.Println("Looping biasa ")
	fmt.Println("")
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

}
