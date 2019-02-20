// Checks if the two provided strings are anagrams of each other
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Need to provide 2 strings as arguments")
		os.Exit(1)
	}

	println(checkAnagram(os.Args[1], os.Args[2]))

}

func checkAnagram(s1, s2 string) bool {
	s1bytes := []byte(s1)
	s2bytes := []byte(s2)

	// check if lengths are qual
	if len(s1bytes) != len(s2bytes) {
		return false
	}

	// check if each rune in s1 appears in s2 the same number of times
	for _, r := range bytes.Split(s1bytes, []byte("")) {
		if bytes.Count(s1bytes, r) != bytes.Count(s2bytes, r) {
			return false
		}
	}

	return true
}
