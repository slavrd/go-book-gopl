package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := []byte("\u0E50\u0E51\u0E52")
	fmt.Println(str)
	fmt.Println(revRunes(str))
}

// reverse a []byte
func revBytes(b []byte) {
	length := len(b)
	for i := 0; i < length/2; i++ {
		b[i], b[length-1-i] = b[length-1-i], b[i]
	}
}

// reverse the characters in a []byte slice
func revRunes(s []byte) []byte {

	// reverse each rune
	for i := 0; i < len(s); {
		_, rlenght := utf8.DecodeRune(s[i:])
		if rlenght > 1 {
			revBytes(s[i : i+rlenght])
		}
		i += rlenght
	}

	// reverse the whole slice
	revBytes(s)
	return s
}
