package split_test

import (
	"strings"
	"testing"
)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		s, sep string
		want   []string
	}{
		{"", "", []string{}},
		{"", ":", []string{""}},
		{"foo", ":", []string{"foo"}},
		{"foo", "", []string{"f", "o", "o"}},
		{"foo:bar", ":", []string{"foo", "bar"}},
		{":foo:", ":", []string{"", "foo", ""}},
		{"foo\u003abar", ":", []string{"foo", "bar"}},
		{"foo:bar", "\u003a", []string{"foo", "bar"}},
		{"foo", "foo", []string{"", ""}},
	}

	for _, test := range tests {
		if got := strings.Split(test.s, test.sep); !equal(got, test.want) {
			t.Errorf("Split(%q, %q) = %v, want %v", test.s, test.sep, got, test.want)
		}
	}
}
