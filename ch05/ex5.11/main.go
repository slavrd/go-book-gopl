// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	depStack := make([]string, 0)
	brokenDeps := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	isDepLoop := false

	visitAll = func(items []string) {

		for _, item := range items {

			// set isDepLoop when begining a new dependency stack
			if len(depStack) == 0 {
				isDepLoop = false
			}

			// Check if current item is already in the current dependency stack
			for i := 0; i < len(depStack); i++ {
				if depStack[i] == item {
					isDepLoop = true
					printDepStackLoop(depStack, item)
					return
				}
			}

			depStack = append(depStack, item) // add to dependency stack

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				if !isDepLoop { // current depStack is marked as broken, skip append to result
					order = append(order, item)
				} else {
					brokenDeps[item] = true
				}
			} else if brokenDeps[item] == true { // if item was seen before and is a known broken dep
				isDepLoop = true
			}

			depStack = depStack[:len(depStack)-1] // pop from dependency stack

		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func printDepStackLoop(s []string, next string) {
	fmt.Println("Dependency loop detected:")
	for _, i := range s {
		fmt.Printf("%s -> ", i)
	}
	fmt.Printf("%s\n", next)
}

//!-main
