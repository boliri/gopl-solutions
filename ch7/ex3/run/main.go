package main

import (
	"fmt"

	"treesort"
)

func main() {
	// Tree topology:
	//         1
	//      /     \
	//    2         3
	//  /   \     /   \
	// 4     5   6     7

	// Level 3 trees
	t31 := treesort.Tree{Value: 4, Left: nil, Right: nil}
	t32 := treesort.Tree{Value: 5, Left: nil, Right: nil}
	t33 := treesort.Tree{Value: 6, Left: nil, Right: nil}
	t34 := treesort.Tree{Value: 7, Left: nil, Right: nil}

	// Level 2 trees
	t21 := treesort.Tree{Value: 2, Left: &t31, Right: &t32}
	t22 := treesort.Tree{Value: 3, Left: &t33, Right: &t34}

	// Top level tree
	root := treesort.Tree{Value: 1, Left: &t21, Right: &t22}

	fmt.Printf("%s\n", &root)
}
