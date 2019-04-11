package intset

import "fmt"

// TestLen demonstrates the Len()
func TestLen() {

	fmt.Printf("\nTesting Len()...\n")

	var s IntSet
	s.Add(1)
	s.Add(12)
	s.Add(798)

	fmt.Printf("The length of set: %s is %d\n", s.String(), s.Len())

	s.Add(15)
	s.Add(1024)

	fmt.Printf("The length of set: %s is %d\n", s.String(), s.Len())
}

// TestRemove demonstrates Remove()
func TestRemove() {
	fmt.Printf("\nTesting Remove()...\n")

	var s IntSet
	s.Add(1)
	s.Add(12)
	s.Add(64)
	s.Add(798)
	s.Add(1024)

	fmt.Printf("Initial set: %s\n", s.String())
	s.Remove(12)
	fmt.Printf("Set after calling Remove(12): %s\n", s.String())
	s.Remove(64)
	fmt.Printf("Set after calling Remove(64): %s\n", s.String())
	s.Remove(1024)
	fmt.Printf("Set after calling Remove(1024): %s\n", s.String())
	s.Remove(1050)
	fmt.Printf("Set after calling Remove(1050): %s\n", s.String())

}

// TestClear demonstrates Clear()
func TestClear() {
	fmt.Printf("\nTesting Clear()...\n")
	var s IntSet
	s.Add(1)
	s.Add(12)
	s.Add(64)
	s.Add(798)
	s.Add(1024)
	fmt.Printf("Initial set: %s\n", s.String())
	s.Clear()
	fmt.Printf("Set after calling Clear(): %s\n", s.String())
}

// TestCopy demonstrates Copy()
func TestCopy() {
	fmt.Printf("\nTesting Copy()...\n")
	var s IntSet
	s.Add(1)
	s.Add(12)
	s.Add(64)
	s.Add(798)
	s.Add(1024)
	fmt.Printf("Initial set1: %s\n", s.String())
	s2 := s.Copy()
	fmt.Printf("Set2 (called Copy on set1): %s\n", s2.String())
	s.Add(2048)
	fmt.Printf("Set1 after calling Add(2048) on it: %s\n", s.String())
	fmt.Printf("Set2 after calling Add(2048) on set1: %s\n", s2.String())
}

// TestAddAll demonstrates AddAll()
func TestAddAll() {
	fmt.Printf("\nTesting AddAll()...\n")
	var s IntSet
	s.Add(1)
	s.Add(12)
	fmt.Printf("Initial set1: %s\n", s.String())
	s.AddAll(1024, 15, 80)
	fmt.Printf("Set after calling AddAll(1024, 15, 80) : %s\n", s.String())
}

// TestIntersectWith demonstrates IntersectWith()
func TestIntersectWith() {
	fmt.Printf("\nTesting IntersectWith()...\n")
	var s, t IntSet
	s.AddAll(1, 12, 145, 758)
	t.AddAll(1, 128, 145)
	fmt.Printf("Initial set1: %s\n", s.String())
	fmt.Printf("Set2: %s\n", t.String())
	s.IntersectWith(&t)
	fmt.Printf("Set1 after intersecting with set2: %s\n", s.String())
	s.IntersectWith(&t)
	fmt.Printf("Set1 after second intersect with set2: %s\n", s.String())
	t.Clear()
	s.IntersectWith(&t)
	fmt.Printf("Set1 after intersect with empty set: %s\n", s.String())
}

// TestDifferenceWith demonstrates DifferenceWith()
func TestDifferenceWith() {
	fmt.Printf("\nTesting DifferenceWith()...\n")
	var s, t IntSet
	s.AddAll(1, 12, 145, 758)
	t.AddAll(1, 128, 145)
	fmt.Printf("Initial set1: %s\n", s.String())
	fmt.Printf("Set2: %s\n", t.String())
	s.DifferenceWith(&t)
	fmt.Printf("Difference of set1 with set2: %s\n", s.String())
}

// TestSymmetricDifference demonstrates SymmetricDifference()
func TestSymmetricDifference() {
	fmt.Printf("\nTesting SymmetricDifference()...\n")
	var s, t IntSet
	s.AddAll(1, 12, 145, 758)
	t.AddAll(1, 128, 145, 1024)
	fmt.Printf("Initial set1: %s\n", s.String())
	fmt.Printf("Set2: %s\n", t.String())
	s.SymmetricDifference(&t)
	fmt.Printf("Symmetric difference of set1 with set2: %s\n", s.String())
}

// TestElems demonstrates Elems()
func TestElems() {
	fmt.Printf("\nTesting Elems()...\n")
	var s IntSet
	s.AddAll(1, 12, 145, 758, 63, 64, 65, 0)
	fmt.Printf("Set1: %s\n", s.String())
	fmt.Printf("Elements of set1 as slice: % #v\n", s.Elems())
}
