// Package bitcompare contains functions for bit comparisons
package bitcompare

// DiffBitCount counts of the same bits in a and b arrays
func DiffBitCount(a [32]byte, b [32]byte) int {

	count := 0
	for index, vala := range a {
		test := vala ^ b[index]
		for test != 0 {
			count += 1 - (int)(test&1)
			test = test >> 1
		}
	}
	return count
}
