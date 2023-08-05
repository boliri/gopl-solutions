// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package findlinks

import (
	"golang.org/x/net/html"
)

//!+
// visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	for _, a := range n.Attr {
		if a.Key == "href" || a.Key == "src" {
			links = append(links, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

//!-