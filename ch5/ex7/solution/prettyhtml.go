package prettyhtml

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var out string

// Prettify prettifies a HTML document by adding tabs in accordance to nesting levels
// and shortening tags that have no children
// Returns a prettified version of the original HTML in string format
func Prettify(root *html.Node) string {
	forEachNode(root, startElement, endElement)
	return out
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if !isPrintable(n) {
		return
	}

	if hasChildren(n) {
		printParentOpenTag(n)
		depth++
	} else {
		if n.Type == html.CommentNode || n.Type == html.TextNode {
			if !hasData(n) {
				depth--
				return
			}

			if n.Type == html.TextNode {
				printText(n)
			} else {
				printComment(n)
			}

			if !hasSiblings(n) {
				depth--
			}
		} else {
			printChildless(n)
		}
	}
}

func endElement(n *html.Node) {
	if !isPrintable(n) || !hasChildren(n) {
		return
	}

	printParentCloseTag(n)

	if !hasSiblings(n) {
		depth--
	}
}

func printParentOpenTag(n *html.Node) {
	attrs := stringify(n.Attr)
	if attrs != "" {
		attrs = " " + attrs
	}

	out += fmt.Sprintf("%*s<%s%s>\n", depth*2, "", n.Data, attrs)
}

func printParentCloseTag(n *html.Node) {
	out += fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
}

func printChildless(n *html.Node) {
	attrs := stringify(n.Attr)
	if attrs != "" {
		attrs = " " + attrs
	}

	out += fmt.Sprintf("%*s<%s%s/>\n", depth*2, "", n.Data, attrs)
}

func printText(n *html.Node) {
	out += fmt.Sprintf("%*s%s\n", depth*2, "", n.Data)
}

func printComment(n *html.Node) {
	out += fmt.Sprintf("%*s<!--%s-->\n", depth*2, "", n.Data)
}

func stringify(attrs []html.Attribute) string {
	var s string
	for _, a := range attrs {
		s += fmt.Sprintf("%s=\"%s\" ", a.Key, a.Val)
	}
	
	return strings.TrimSpace(s)
}

func hasChildren(n *html.Node) bool {
	return n.FirstChild != nil
}

func hasSiblings(n *html.Node) bool {
	return n.NextSibling != nil
}

func hasData(n *html.Node) bool {
	return strings.TrimSpace(n.Data) != ""
}

func isPrintable(n *html.Node) bool {
	return n.Type == html.ElementNode || n.Type == html.CommentNode || n.Type == html.TextNode
}