package main

import "fmt"

func main() {

	// slice sebenarnya bukanlah suatu data
	// melainkan seperti sebuah pointer yang mana menunjukkan
	// data dari posisi ke posisi
	// ia mereferensikan data array sebelumnya

	names := [5]string{
		"Reyhan",
		"Jovie",
		"Dwiputra",
		"NooBee",
		"Achimonchi",
	}

	fmt.Println(names)

	// slice berguna untuk memotong / mengambil array berdasarkan index
	slice1 := names[1:4] // outputnya adalah Jovie, Dwiputra, NooBee
	slice2 := names[:3]  // outputnya adalah Reyhan, Jovie, Dwiputra

	fmt.Println("slice 1 ", slice1)
	fmt.Println("slice 2 ", slice2)

	fmt.Println()
	names[2] = "Programming"
	fmt.Println(names)
	fmt.Println("slice 1 ", slice1)
	fmt.Println("slice 2 ", slice2)

	// ini untuk membuat slice
	// data pada slice, adalah sebuah array yang berada
	// didalam slice
	fmt.Println()

	slice := make([]string, 3)
	slice[0] = "Reyhan"
	slice[1] = "Jovie"
	slice[2] = "Dwiputra"

	fmt.Println(slice)

	// karena slice mempunyai array di dalamnya, maka bisa melakukan
	// slice di dalam slice
	slice4 := slice[:]

	fmt.Println()
	fmt.Println("slice  ", slice)
	fmt.Println("slice4 ", slice4)

	// untuk menambahkan data pada slice, bisa menggunakan append
	// namun perlu diingat, bahwa hal ini akan menyebabkan
	// membuat slice dan array baru

	slice5 := append(slice, "orang2", "orang3")
	fmt.Println()
	fmt.Println("slice  ", slice)
	fmt.Println("slice5 ", slice5)

	// lalu, jika dilakukan perubahan pada nilai di slice
	slice[1] = "Ubah si Jovie"
	fmt.Println()
	fmt.Println("slice  ", slice)
	fmt.Println("slice5 ", slice5)
	// nah, dapat dilihat bahwa slice5 tidak terpengaruh dengan perubahan pada
	// slice. karena, si slice5 melakukan copy terhadap si slice, bukan
	// mereferences nya

	// ada method copy untuk mengcopy seluruh isi slice
	// method ini juga akan melakukan copy sebesar ukuran si
	// destination slice
	// jika ukuran destination slice = 4, dan source slice = 10,
	// maka data yang tercopy cuma dari data pertama - data keempat (indeks 0 - 3)
	slice6 := make([]string, 2)
	copy(slice6, slice)
	fmt.Println()
	fmt.Println("slice  ", slice)
	fmt.Println("slice6 ", slice6)

}
