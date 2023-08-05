// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import "fmt"

// !+

// NOTE: Struct fields were made public to make them fillable from the main package
type Tree struct {
	Value       int
	Left, Right *Tree
}

func (t *Tree) String() string {
	str := ""

	currentLevelTrees := []*Tree{t}
	var nextLevelTrees []*Tree

	for len(currentLevelTrees) > 0 {
		for _, tree := range currentLevelTrees {
			str += fmt.Sprintf("%d ", tree.Value)

			if tree.Left != nil {
				nextLevelTrees = append(nextLevelTrees, tree.Left)
			}

			if tree.Right != nil {
				nextLevelTrees = append(nextLevelTrees, tree.Right)
			}
		}

		str += "\n"

		currentLevelTrees = nextLevelTrees
		nextLevelTrees = nil
	}

	return str
}

//!-
