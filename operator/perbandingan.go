package main

import "fmt"

func main() {
	a := 1
	b := 2

	// perbandingan di golang sama dengan perbandingan
	// di bahasa lain
	// == , <= , >= , < , > , !=

	// result := a == b   	=> false
	// result := a <= b		=> true
	// result := a >= b		=> false
	// result := a < b		=> true
	// result := a > b		=> false
	// result := a != b		=> true

	result := a == b
	fmt.Println(result)
}
