// Charcount the type of characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	invalid := 0                   // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsSpace(r):
			counts["spaces"]++
		case unicode.IsNumber(r):
			counts["numbers"]++
		case unicode.IsLetter(r):
			counts["letters"]++
		case unicode.IsPunct(r):
			counts["punctuation"]++
		case unicode.IsMark(r):
			counts["marks"]++
		default:
			counts["unknown"]++
		}
	}

	fmt.Println()
	for k := range counts {
		fmt.Printf("%s:\t%v\n", k, counts[k])
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
