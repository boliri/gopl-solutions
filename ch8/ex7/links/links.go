package links

import (
	"net/url"

	"golang.org/x/net/html"
)

// ExtractLinksWrapper returns a function that accepts an HTML node as input so forEachNode can use it
//
// this inner function extracts links from <a> elements with href attributes
func ExtractLinksWrapper(links *[]string, baseUrl *url.URL) func(n *html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := baseUrl.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				*links = append(*links, link.String())
			}
		}
	}
}

// TransformHrefsWrapper returns a function that accepts an HTML node as input so forEachNode can use it
//
// this inner function transforms href values according to the transformFn function
func TransformHrefsWrapper(transformFn func(string) string, baseUrl *url.URL) func(n *html.Node) {
	return func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := baseUrl.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}

				n.Attr[i].Val = transformFn(link.String())
			}
		}
	}
}

// ProcessDocument deals with an HTML document according to the pre and post functions
func ProcessDocument(doc *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(doc)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		ProcessDocument(c, pre, post)
	}
	if post != nil {
		post(doc)
	}
}
