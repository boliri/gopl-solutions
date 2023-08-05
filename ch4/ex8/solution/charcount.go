// Charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func Count(in *bufio.Reader) {
	letterCounts := make(map[rune]int)    // counts of Unicode letter characters
	digitCounts := make(map[rune]int)    // counts of Unicode digit characters
	symbolCounts := make(map[rune]int)    // counts of Unicode symbol characters
	markCounts := make(map[rune]int)    // counts of Unicode mark characters
	punctCounts := make(map[rune]int)    // counts of Unicode punctuation characters

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letterCounts[r]++
		} else if unicode.IsDigit(r) {
			digitCounts[r]++
		} else if unicode.IsSymbol(r) {
			symbolCounts[r]++
		} else if unicode.IsMark(r) {
			markCounts[r]++
		} else if unicode.IsPunct(r) {
			punctCounts[r]++
		}
		utflen[n]++
	}
	fmt.Println("\nletters")
	fmt.Printf("rune\tcount\n")
	for c, n := range letterCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("\ndigits")
	fmt.Printf("rune\tcount\n")
	for c, n := range digitCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("\nsymbols")
	fmt.Printf("rune\tcount\n")
	for c, n := range symbolCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("\nmarks")
	fmt.Printf("rune\tcount\n")
	for c, n := range markCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("\npuncts")
	fmt.Printf("rune\tcount\n")
	for c, n := range punctCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("\nutf-8 per size in bytes")
	fmt.Print("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}