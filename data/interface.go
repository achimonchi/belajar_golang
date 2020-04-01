package main

import (
	"fmt"
	"math"
)

// deklarasi kan interface
// pada interface, diletakkan function
// yang akan dimiliki oleh si struct
type Shape interface {
	// deklarasi method pada interface
	Area() float64
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

// agar struct dapat mengimplementasi Shape,
// maka harus mengimplementasi seluruh method pada
// interface Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// hal ini menyebabkan siapapun yang telah mengimplementasi
// interface Shape, boleh menggunakan method ini
func getArea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{
		Radius: 14.0,
		X:      30,
		Y:      20,
	}

	rectangle := Rectangle{
		Height: 30,
		Width:  20,
	}

	fmt.Printf("Area dari circle = %f", getArea(circle))
	fmt.Println()
	fmt.Printf("Area dari rectangle = %f", getArea(rectangle))
	fmt.Println()
}
