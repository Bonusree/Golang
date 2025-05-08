package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	x := (x1 - x2)
	y := (y1 - y2)
	return math.Sqrt(x*x + y*y)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)
	return l * w
}
func main() {
	r := Rectangle{2, 4, 2, 4}
	fmt.Println(r.area())
}
