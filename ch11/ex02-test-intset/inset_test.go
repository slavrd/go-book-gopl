package intset

import (
	"reflect"
	"testing"
)

func TestHas(t *testing.T) {

	// set (0,1,64,70,200) as a bit vector
	testSet := IntSet{
		words: []uint64{3, (1 + (1 << 6)), 0, (1 << 8)},
	}

	// contents of the testSet as integers, to use for display
	visualTestSet := []int{0, 1, 64, 70, 200}

	tcs := []struct {
		tvalue int
		result bool
	}{
		{0, true},
		{1, true},
		{64, true},
		{70, true},
		{200, true},
		{2, false},
		{130, false},
		{500, false},
	}

	for _, tc := range tcs {
		if r := testSet.Has(tc.tvalue); r != tc.result {
			t.Errorf("set: %v, bit vector words: %v got: %v for value: %v", visualTestSet, testSet, r, tc.tvalue)
		}
	}
}

func TestAdd(t *testing.T) {

	// set (0,1,64,70,200) as a bit vector
	testSet := IntSet{
		words: []uint64{3, (1 + (1 << 6)), 0, (1 << 8)},
	}

	tests := []struct {
		set      IntSet
		addition int
		want     IntSet
	}{
		{
			set:      testSet,
			addition: 2,
			want:     IntSet{words: []uint64{7, (1 + (1 << 6)), 0, (1 << 8)}},
		},
		{
			set:      testSet,
			addition: 128,
			want:     IntSet{words: []uint64{7, (1 + (1 << 6)), 1, (1 << 8)}},
		},
		{
			set:      testSet,
			addition: 321,
			want:     IntSet{words: []uint64{7, (1 + (1 << 6)), 1, (1 << 8), 0, 2}},
		},
	}

	for _, test := range tests {
		test.set.Add(test.addition)
		if !reflect.DeepEqual(test.want, test.set) {
			t.Errorf("add: %v to bit vector set: %v got: %v wanted: %v",
				test.addition, testSet, test.set, test.want)
		}
	}
}

func TestUnionWith(t *testing.T) {

	// expected set (0,1,2,64,65,70,130,200,320,321) as a bit vector
	rSet := IntSet{
		words: []uint64{7, (1 + 2 + (1 << 6)), (1 << 2), (1 << 8), 0, 3},
	}

	tests := []struct {
		set1    IntSet
		set2    IntSet
		wantSet IntSet
	}{
		{
			// set (0,1,64,70,200) as a bit vector
			set1: IntSet{words: []uint64{3, (1 + (1 << 6)), 0, (1 << 8)}},
			// set (0,2,64,65,130,200,320,321) as a bit vector
			set2:    IntSet{words: []uint64{5, 3, 4, (1 << 8), 0, 3}},
			wantSet: rSet,
		},
		{
			// set (0,2,64,65,130,200,320,321) as a bit vector
			set1: IntSet{words: []uint64{5, 3, 4, (1 << 8), 0, 3}},
			// set (0,1,64,70,200) as a bit vector
			set2:    IntSet{words: []uint64{3, (1 + (1 << 6)), 0, (1 << 8)}},
			wantSet: rSet,
		},
	}

	for _, test := range tests {

		// create a copy of test.set1 on which to call UnionWith so that test.set1 is preserved
		var s IntSet
		s.words = make([]uint64, len(test.set1.words))
		copy(s.words, test.set1.words)

		s.UnionWith(&test.set2)

		if !reflect.DeepEqual(s, test.wantSet) {
			t.Errorf("set1: %v union with set2: %v got: %v want: %v",
				test.set1, test.set2, s, test.wantSet)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		set  IntSet
		want string
	}{
		// Testcase empty set
		{
			set:  IntSet{words: []uint64{}},
			want: "{}",
		},
		// Testcase set {0 1 65 200}
		{
			set:  IntSet{words: []uint64{3, 2, 0, (1 << 8)}},
			want: "{0 1 65 200}",
		},
	}

	for _, test := range tests {
		r := test.set.String()
		if r != test.want {
			t.Errorf("set: %v to string got: %q want: %q",
				test.set, r, test.want)
		}
	}
}
