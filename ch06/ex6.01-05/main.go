package main

import (
	"github.com/go-book-gopl/ch06/ex6.01-05/intset"
)

func main() {
	intset.TestLen()
	intset.TestRemove()
	intset.TestClear()
	intset.TestCopy()
	intset.TestAddAll()
	intset.TestIntersectWith()
	intset.TestDifferenceWith()
	intset.TestSymmetricDifference()
	intset.TestElems()
}
