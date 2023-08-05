package findelems

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return visit([]*html.Node{}, doc, name...)
}

// visit appends to elems each element found in n with a matching tag name
// and returns the result.
func visit(elems []*html.Node, n *html.Node, name ...string) []*html.Node {
	if n.Type == html.ElementNode && contains(name, n.Data) {
		elems = append(elems, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elems = visit(elems, c, name...)
	}

	return elems
}

// contains checks if a string is in a slice
func contains(list []string, s string) bool {
	for _, elem := range list {
		if elem == s {
			return true
		}
	}
	return false
}