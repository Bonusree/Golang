package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	} //infinite loop
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
