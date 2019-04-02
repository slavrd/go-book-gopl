package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Minimum value via min: %v\n", min(1, 2, 3, -10, -4, 5, 0, 7))
	fmt.Printf("Minimum value via min1: %v\n", min(1, 2, 3, -10, -4, 5, 0, 7))
	fmt.Printf("Maximum value via max: %v\n", max(1, 2, 3, -10, -4, 5, 0, 7))
	fmt.Printf("Maximum value via max1: %v\n", max(1, 2, 3, -10, -4, 5, 0, 7))

}

func min(vals ...int) int {

	if len(vals) == 0 {
		return math.MinInt32
	}

	min := vals[0]

	for _, val := range vals[1:] {
		if min > val {
			min = val
		}
	}

	return min
}

func min1(i int, vals ...int) int {

	min := i

	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}

func max(vals ...int) int {

	if len(vals) == 0 {
		return math.MaxInt32
	}

	max := vals[0]

	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}

	return max
}

func max1(i int, vals ...int) int {

	max := i

	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}
