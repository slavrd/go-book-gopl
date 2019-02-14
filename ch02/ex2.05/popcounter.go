package main

import (
	"fmt"
	"os"
	"strconv"

	popcount "github.com/go-book-gopl/ch02/ex2.05/popcount"
)

// Main program
func main() {
	for _, arg := range os.Args[1:] {
		value, _ := strconv.Atoi(arg)
		fmt.Printf("PopCount of %s is %v\n", arg, popcount.PopCount(uint64(value)))
	}
}
