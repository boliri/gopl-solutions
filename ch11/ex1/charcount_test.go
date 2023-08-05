package charcount

import (
	"bytes"
	"io"
	"testing"
	"unicode/utf8"
)

func compareCounts(c1, c2 map[rune]int) bool {
	for k, v1 := range c1 {
		if v2, ok := c2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input       io.Reader
		wantCount   map[rune]int
		wantUtfLen  [utf8.UTFMax + 1]int
		wantInvalid int
	}{
		{
			bytes.NewReader([]byte("Hello, world")),
			map[rune]int{'H': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, ' ': 1, 'w': 1, 'r': 1, 'd': 1},
			[utf8.UTFMax + 1]int{0, 12, 0, 0, 0},
			0,
		},
		{
			bytes.NewReader([]byte("Hello, 世界")),
			map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1, ',': 1, ' ': 1, '世': 1, '界': 1},
			[utf8.UTFMax + 1]int{0, 7, 0, 2, 0},
			0,
		},
		// I don't know how to force a string to have invalid unicode characters without the
		// compiler complaining about illegal chars, so I'll leave this test out
		// {
		// 	bytes.NewReader([]byte("Hello, \uFFFG")),
		// 	map[rune]int{'H': 1, 'e': 1, 'l': 2, 'o': 1, ',': 1, ' ': 1},
		// 	[utf8.UTFMax + 1]int{0, 7, 0, 0, 0},
		// 	1,
		// },
	}

	for _, test := range tests {
		c := NewCounter(test.input)
		c.Count()
		gotCount, gotUtfLen, gotInvalid := c.Results()

		if !compareCounts(gotCount, test.wantCount) {
			t.Errorf("Got count = %v, want %v", gotCount, test.wantCount)
		}
		if gotUtfLen != test.wantUtfLen {
			t.Errorf("Got utflen = %v, want %v", gotUtfLen, test.wantUtfLen)
		}
		if gotInvalid != test.wantInvalid {
			t.Errorf("Got invalid = %v, want %v", gotInvalid, test.wantInvalid)
		}
	}
}
