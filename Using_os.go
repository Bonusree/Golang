package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Exiting the program...")
	os.Exit(1)
	fmt.Println("This will NOT be printed") // this line is skipped
}
