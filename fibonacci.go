package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci sequence:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fib(i))
	}
}
