package popcount

// PopCount returns the population count (number of set bits) of x > 0
func PopCount(x uint64) int {
	test := x
	count := 0
	for test != 0 {
		count += (int)(test & 1)
		test = test >> 1
	}
	return count
}
