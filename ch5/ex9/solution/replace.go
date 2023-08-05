package replace

import (
	"fmt"
	"strings"
)

var replaceable string = "$foo"
var seed string = "foo"


// Expand replaces all occurences of the substring held by "replaceable" with the outcome
// of the function f, which takes "seed" to generate the replacement string
func Expand(s string, f MutateFn) string {
	replacement := f(seed)
	return strings.ReplaceAll(s, replaceable, replacement)
}


// Type for functions that take a string and mutate it using a specific method
type MutateFn func(string) string

// Caesar takes a string and returns the same string but Caesar-cyphered (shift=1)
func Caesar(s string) string {
	var res string
	for _, c := range s {
		res += string(c + 1)
	}

	return res
}

// Hex takes a string and returns the same string in hexadecimal format
func Hex(s string) string {
	return fmt.Sprintf("%X", s)
}

// Noop takes a string and returns it as is
func Noop(s string) string {
	return s
}
