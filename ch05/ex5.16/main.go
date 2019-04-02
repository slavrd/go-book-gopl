package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("Result of varJoin: %s\n", varJoin("+", "aa", "bb", "cc"))

}

func varJoin(sep string, s ...string) string {

	var slice []string

	for _, str := range s {
		slice = append(slice, str)
	}

	if len(slice) > 0 {
		return strings.Join(slice, sep)
	} else {
		return ""
	}

}
