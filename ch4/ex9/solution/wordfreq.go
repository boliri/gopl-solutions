// wordfreq computes counts of words in a given file

package wordfreq

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Count(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Something happened while trying to open %s: %s\n", filename, err)
		os.Exit(1)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	wordCounts := make(map[string]int)
	for s.Scan() {
		wordCounts[strings.ToLower(s.Text())]++
	}

	fmt.Printf("word counts for %s\n", filename)
	for w, n := range wordCounts {
		fmt.Printf("%s\t\t%d\n", w, n)
	}
}