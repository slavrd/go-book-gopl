// Rotates a slice x elements to the left
package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions
	rotate(s, 2)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func rotate(s []int, e int) {
	result := make([]int, len(s), cap(s))
	copy(result, s[e:])
	copy(result[len(s[e:]):], s[:e])
	copy(s, result)
}
