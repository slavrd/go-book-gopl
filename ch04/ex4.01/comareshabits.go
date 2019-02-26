package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/go-book-gopl/ch04/ex4.01/bitcompare"
)

func main() {

	// Check the passed arguments
	if len(os.Args) != 3 {
		fmt.Println("Function takes 2 arguments for which to compare sha256 bits")
		os.Exit(1)
	}

	a := sha256.Sum256([]byte(os.Args[1]))
	b := sha256.Sum256([]byte(os.Args[2]))

	fmt.Printf("Sha256 arg1: %x\nSha256 arg2: %x\n", a, b)

	diffBits := bitcompare.DiffBitCount(a, b)
	fmt.Printf("The number of different bits is: %v\n", diffBits)
}
