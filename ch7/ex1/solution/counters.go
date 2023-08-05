package counters

import (
	"bufio"
	"bytes"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bufio.NewReader(bytes.NewReader(p)))
	s.Split(bufio.ScanWords)

	var words int
	for s.Scan() {
		words += 1
	}

	*c = WordCounter(words)
	return words, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bufio.NewReader(bytes.NewReader(p)))
	s.Split(bufio.ScanLines)

	var lines int
	for s.Scan() {
		lines += 1
	}

	*c = LineCounter(lines)
	return lines, nil
}
