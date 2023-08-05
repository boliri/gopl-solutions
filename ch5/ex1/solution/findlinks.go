// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package findlinks

import (
	"golang.org/x/net/html"
)

//!+
// visit appends to links each link found in n and returns the result.
func Visit(links []string, unvisited []*html.Node) []string {
	n := unvisited[0]
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		unvisited = append(unvisited, n.FirstChild)
	}
	if n.NextSibling != nil {
		unvisited = append(unvisited, n.NextSibling)
	}

	remaining := unvisited[1:]
	if len(remaining) == 0 {
		return links
	} else {
		return Visit(links, remaining)
	}
}

//!-