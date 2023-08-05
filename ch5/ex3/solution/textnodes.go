// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package textnodes

import (
	"unicode"

	"golang.org/x/net/html"
)

//!+
// find takes a pointer to a slice of strings to update the underlying array with all the text nodes in a HTML document
func Find(t *[]string, n *html.Node) {
	if n.Type == html.TextNode && !unicode.IsSpace(rune(n.Data[0])) {
		*t = append(*t, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Find(t, c)
	}
}

//!-