package main

import "fmt"

func main() {
	var x uint64 = 259
	fmt.Println(x)
	fmt.Println(byte(x))

	fmt.Println(byte(5 & 1))
}
