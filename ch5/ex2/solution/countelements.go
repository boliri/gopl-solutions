// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package countelements

import (
	"golang.org/x/net/html"
)

//!+
// count populates a mapping with HTML names and the number of times they appear in a HTML document.
func Count(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Count(counts, c)
	}
}

//!-