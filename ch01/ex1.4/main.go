// prints the count and text of lines that
// appear more than once in the named input files.
// Also prints the filename in which they appear
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string][]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			addFileName(filename, files, line)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\nInFiles:\n", n, line)
			for _, f := range files[line] {
				fmt.Printf("%s\n", f)
			}
			fmt.Println()
		}
	}
}

func addFileName(file string, fslice map[string][]string, line string) {
	notexists := true
	for _, f := range fslice[line] {
		if f == file {
			notexists = false
		}
	}
	if notexists {
		fslice[line] = append(fslice[line], file)
	}
}
