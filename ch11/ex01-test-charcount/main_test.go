package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {

	type CCTestCase struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}

	tc := []CCTestCase{
		{
			input:   "aB1!@@@世a",
			counts:  map[rune]int{'a': 2, 'B': 1, '1': 1, '!': 1, '@': 3, '世': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 8, 0, 1, 0},
			invalid: 0,
		},
	}

	for _, test := range tc {
		counts, utflen, invalid := charCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) {
			t.Errorf("runes got: %v expected: %v from input: %q", counts, test.counts, test.input)
		}
		if !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("UTF lengths got: %v expected: %v from input: %q", utflen, test.utflen, test.input)
		}
		if invalid != test.invalid {
			t.Errorf("invalids got: %v expected: %v from input: %q", invalid, test.invalid, test.input)
		}
	}

}
