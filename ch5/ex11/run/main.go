// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

package main

import (
	"fmt"

	"toposort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": 		{"data structures"},
	"calculus":   		{"linear algebra"},
	"linear algebra":	{"calculus"},
	
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
}

//!-table

func main() {
	for i, course := range toposort.TopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
