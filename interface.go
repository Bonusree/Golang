package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return 4 * s.side
}
func main() {
	var shape Shape
	shape = Circle{5}
	fmt.Println("Area of circle: ", shape.Area())
	fmt.Println("Perimeter of circle: ", shape.Perimeter())
	shape = Square{5}
	fmt.Println("Area of square: ", shape.Area())
	fmt.Println("Perimeter of square: ", shape.Perimeter())

}
