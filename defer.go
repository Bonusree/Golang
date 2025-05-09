package main

import "fmt"

func main() {
	//file, _ := os.Create("file.txt")
	//defer file.Close()
	//_, _ = io.WriteString(file, "writing text")
	// for defer print value in last moment of execution
	defer fmt.Println("hello")
	fmt.Println("bye")
}
