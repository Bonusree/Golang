package main

import (
	"fmt"
)

func handlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("a is not nil", a)
	}
}
func main() {
	fmt.Println("hello")
	runPanic() //here will be stopped program
	fmt.Println("bye")
}
func runPanic() {
	defer handlePanic()
	panic("panic")
}
