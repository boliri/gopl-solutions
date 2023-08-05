// Charcount computes counts of Unicode characters.
package charcount

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// Counter holds a buffer and count-related statistics of such buffer
type counter struct {
	r       io.Reader
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func NewCounter(r io.Reader) *counter {
	return &counter{r: r, counts: make(map[rune]int)}
}

// Count counts Unicode characters from the counter's underlying buffer
func (c *counter) Count() {
	s := bufio.NewScanner(c.r)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		r, size := utf8.DecodeRune(s.Bytes())
		if r == utf8.RuneError && size == 1 {
			c.invalid++
			continue
		}

		c.counts[r]++
		c.utflen[size]++
	}
}

// Results returns the stats of a counter
func (c *counter) Results() (counts map[rune]int, utflen [utf8.UTFMax + 1]int, invalid int) {
	return c.counts, c.utflen, c.invalid
}
