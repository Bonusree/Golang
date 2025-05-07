package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " even ")
		} else if i%3 == 1 {
			fmt.Println(i, " divided by 3")
		} else {
			fmt.Println(i, " odd and not divided by 3")
		}

	}
}
