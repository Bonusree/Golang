package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func square(x *float64) {
	*x = *x * *x
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(&x) // x is 0
	fmt.Println(x)
	n := 1.3
	square(&n)
	fmt.Println(n)
	y := 1
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
