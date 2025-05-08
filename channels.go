package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 5)

	// Anonymous goroutine
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "geeks"
		mychnl <- "ggg"
		mychnl <- "gggg"
		mychnl <- "gggg2"
		mychnl <- "GeeksforGeeks"
		close(mychnl)
	}()

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
		fmt.Println("Length of the channel is: ", len(mychnl))
		fmt.Println("Capacity of the channel is: ", cap(mychnl)) //input can be taken more than cap
	}
}
