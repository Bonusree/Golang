package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1+1 = ", 1.0+1.0)

	//string
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) //return byte
	fmt.Println("Hello " + "World")

	//Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//problems
	fmt.Println("Problems: ")
	fmt.Println(int64(math.Pow(10, 8) - 1))
	fmt.Println(321325 * 424521)
	fmt.Println((true && false) || (false && true) || !(false && false))

}
