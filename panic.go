package main

import "fmt"

func main() {
	fmt.Println("hello")
	runPanic() //here will be stopped program
	fmt.Println("bye")
}
func runPanic() {
	panic("panic")
}
