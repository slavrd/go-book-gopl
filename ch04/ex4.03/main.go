// Rev reverses an array
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array
}

func reverse(arr *[6]int) {
	for i := 0; i < len(arr)/2; i++ {
		endIndex := len(arr) - i - 1
		arr[i], arr[endIndex] = arr[endIndex], arr[i]
	}
}

//!-rev
