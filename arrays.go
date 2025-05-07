package main

import "fmt"

func main() {
	var x [5]float64

	// Taking input from user
	fmt.Println("Enter 5 float numbers:")
	for i := 0; i < len(x); i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&x[i])
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(x[0])
	// Printing average
	fmt.Println("Average:", total/float64(len(x)))
	slice_x := x[0:2]
	for _, value := range slice_x {
		fmt.Println(value)
	}
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	fmt.Println(slice1[2:4]) //runtime error

}
