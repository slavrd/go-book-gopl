// rmdups removes the adjacent duplicates in a []string
package main

import "fmt"

func main() {
	s := []string{"item", "item1", "item1", "item1", "item2", "item", "item", "item", "item3"}
	fmt.Println(rmdups(s))
}

// rmdups removes all adjasent duplicates of the slice. Modifies the underlying array
func rmdups(s []string) []string {
	for i := 1; i < len(s); {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			continue
		}
		i++
	}
	return s
}
