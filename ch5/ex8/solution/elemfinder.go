package elemfinder

import (
	"golang.org/x/net/html"
)

var found bool

// Find takes the root element of a HTML document and a string representing an element ID
// Returns a boolean indicating whether an element with the passed ID exists in the document or not
func Find(root *html.Node, id string) bool {
	lookupFn := curriedElementById(root, id)

	forEachNode(root, lookupFn, nil)
	return found
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) *html.Node) {
	if pre != nil && pre(n) != nil {
		found = true
	}

	if found { return }

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func elementById(n *html.Node, id string) *html.Node {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return n
		}
	}

	return nil
}

func curriedElementById(n *html.Node, id string) func(n *html.Node) *html.Node {
	return func(n *html.Node) *html.Node {
		return elementById(n, id)
	}
}