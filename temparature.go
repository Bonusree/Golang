package main

import (
	"fmt"
)

func main() {
	var c, f float64

	fmt.Print("Enter farenhit temperature: ")
	fmt.Scanln(&f)

	c = (f - 32) * 5 / 9

	fmt.Println("celcius temperature =", c)
}
