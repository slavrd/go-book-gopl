package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("  1   2  3  4    ")
	fmt.Println(squashSpace(str))
}

func squashSpace(s []byte) []byte {
	runes := bytes.Runes(s)
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsSpace(runes[i]) {
			runes[i] = rune(32)
			for i+1 < len(runes) && unicode.IsSpace(runes[i+1]) {
				copy(runes[i+1:], runes[i+2:])
				runes = runes[:len(runes)-1]
			}
		}
	}
	return []byte(string(runes))
}
