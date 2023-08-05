package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
    filesWithDupes := make(map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesWithDupes)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesWithDupes)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	
	if len(filesWithDupes) > 0 {
		fmt.Println("List of files with duplicates:")
		for filename, _ := range filesWithDupes {
			fmt.Println(filename)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesWithDupes map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if counts[input.Text()] > 1 {
			filesWithDupes[f.Name()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
