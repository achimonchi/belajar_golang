package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Div(x int, y int) (int, error) {
	// lakukan pengecekan, apabila y adalah 0
	// maka munculkan error
	if y == 0 {
		return 0, errors.New("Tidak bisa membagi dengan 0")
	}

	// jika y != 0, maka errornya akan diset menjadi nil / kosong
	result := x / y
	return result, nil
}

func main() {
	// pada Golang, tidak ada error exception
	// error adalah suatu nilai / memiliki nilai
	iStr := "5a"

	i, err := strconv.Atoi(iStr)

	// karena err memiliki nilai, maka untuk handling bisa dilakukan menggunakan
	// if statement
	// if err != nil, artinya variable err mempunya nilai, maka error
	// namun jika err == nil, maka program berhasil
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	} else {
		fmt.Println("Nilai i adalah", i)
	}

	r, err := Div(18, 3)
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	} else {
		fmt.Println("Hasilnya adalah", r)
	}
}
