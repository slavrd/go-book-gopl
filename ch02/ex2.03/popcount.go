package main

import "fmt"
import "os"
import "strconv"

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	result := 0
	for count := uint(0); count <= 7; count++ {
		result += int(pc[byte(x>>(count*8))])
	}
	return result

}

// Main program
func main() {
	for _, arg := range os.Args[1:] {
		value, _ := strconv.Atoi(arg)
		fmt.Printf("PopCount of %s is %v\n", arg, PopCount(uint64(value)))
	}
}
