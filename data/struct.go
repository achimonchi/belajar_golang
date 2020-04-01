package main

import "fmt"

type Vector struct {
	x int
	y int
}

type Player struct {
	ID   int
	Name string
}

type Person struct {
	ID   int
	Name string
	Age  int
}

// struct sebagai parameter pada suatu function
func printPerson(p Person) {
	fmt.Println(p)
	fmt.Println("ID =", p.ID)
	fmt.Println("Name =", p.Name)
	fmt.Println("Age =", p.Age)
}

// membuat fungsi sebagai member dari struct
// () sebelum GetID disebut Receiver
func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	// struct berupa kumpulan data
	// struct dapat memiliki tipe data yang berbeda
	// bisa menggunakan static declaration
	var v Vector
	v.x = 10
	v.y = 20

	fmt.Println(v)
	fmt.Println("Nilai x =", v.x)
	fmt.Println("Nilai y =", v.y)

	fmt.Println()
	// ataupun dinamic declaration
	player1 := Player{ID: 1, Name: "Jovie"}
	fmt.Println(player1)
	fmt.Println("ID =", player1.ID)
	fmt.Println("Name =", player1.Name)

	fmt.Println()

	person1 := Person{
		ID:   1,
		Name: "Reyhan",
		Age:  21,
	}
	// struct sebagai parameter suatu function
	printPerson(person1)

	fmt.Println()
	// cara memanggil function dari suatu struct adalah
	// sebagai berikut.
	// melakukan pointer terhadap struct
	person2 := &Person{
		ID:   2,
		Name: "Jovie",
		Age:  20,
	}

	// ini adalah cara memanggil functionnya
	fmt.Println(person2)
	fmt.Println("ID =", person2.GetID())
	fmt.Println("Name =", person2.GetName())
	fmt.Println("Age =", person2.GetAge())

	fmt.Println()

	person2.ChangeName("Dwiputra")
	fmt.Println("Dilakukan perubahan nama pada person2")
	fmt.Println("ID =", person2.GetID())
	fmt.Println("Name =", person2.GetName())
	fmt.Println("Age =", person2.GetAge())
}
