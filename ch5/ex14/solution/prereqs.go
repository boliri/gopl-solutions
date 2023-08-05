package solution

import (
	"fmt"
)

var Prereqs map[string][]string = map[string][]string{
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
}

func PrereqsSummary(course string) []string {
	p, ok := Prereqs[course]
	if !ok {
		fmt.Printf("No prereqs found for %s\n\n", course)
	} else {
		fmt.Printf("Prereqs for %s:\n", course)
		for _, prereq := range p {
			fmt.Printf("\t%s\n", prereq)
		}
		fmt.Println()
	}

	return p
}
