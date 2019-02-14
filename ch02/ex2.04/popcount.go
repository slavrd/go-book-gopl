package main

import "fmt"
import "os"
import "strconv"

// PopCount returns the population count (number of set bits) of x > 0
func PopCount(x uint64) int {
	test := x
	count := 0
	for test != 0 {
		count += (int)(test & 1)
		test = test >> 1
	}
	return count
}

// Main program
func main() {
	for _, arg := range os.Args[1:] {
		value, _ := strconv.Atoi(arg)
		fmt.Printf("PopCount of %s is %v\n", arg, PopCount(uint64(value)))
	}
}
