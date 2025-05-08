package main

import (
	"fmt"
	"hash/crc32"
	"hash/fnv"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
	h1 := fnv.New128()
	h1.Write([]byte("test"))
	v1 := h1.Sum() //error: not enough arguments in call to h1.Sum
	fmt.Println(v1)
}
