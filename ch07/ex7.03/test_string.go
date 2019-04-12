package treesort

import "fmt"

func TestString() {

	// Create tree
	t := new(tree)
	values := [...]int{1, 7, 18, 16, 100, 50, 70, -10}

	for _, value := range values {
		t = add(t, value)
	}

	fmt.Println(t.String())

}
