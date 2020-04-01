package main

import "fmt"

// membuat fungsi sederhana
func test() {
	fmt.Println("Fungsi test dipanggil !")
}

func plus(num1 int, num2 int) int {
	return num1 + num2
}

var name string
var age int

func setName(newName string) {
	name = newName
}

func getName() string {
	return name
}

func getBio(age int, name string, job string) {
	fmt.Println(age, name, job)
}

func setAge(newAge int) {
	age = newAge
}

func getAge() int {
	return age
}

// ngereturn 2 value pada function
func getAll() (string, int) {
	return name, age
}

// memberi nama nilai yang di return
func predict(age int, name string) (ageIn10 int, bio string) {
	ageIn10 = age + 10
	bio = "Nama saya adalah " + name
	return
}

func main() {
	test()

	fmt.Println()
	fmt.Println(plus(3, 5))

	fmt.Println()
	setName("Jovie")
	getBio(21, getName(), "Programmer")

	fmt.Println()
	setName("Reyhan")
	setAge(20)

	// untuk mengambil 2 value yang direturn oleh function
	// getAll
	name, age := getAll()

	fmt.Println("Nama", name, "berumur", age)

	fmt.Println()
	age1, name1 := predict(21, "Dwiputra")
	fmt.Println(name1)
	fmt.Println("10 tahun lagu, saya berumur", age1)
}
