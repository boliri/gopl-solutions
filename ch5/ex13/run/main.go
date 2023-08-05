package main

import (
	"os"

	"solution"
)

func main() {
	// Crawl the web breadth-first, starting from the command-line arguments.
	solution.BreadthFirst(solution.Crawl, []string{os.Args[1]})
}
