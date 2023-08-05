package links

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var Done = make(chan struct{})

// requestCtx controls cancellations of all HTTP requests made through the Extract method
type requestCtx int

func (ctx *requestCtx) Deadline() (deadline time.Time, ok bool) { return }
func (ctx *requestCtx) Done() <-chan struct{}                   { return Done }
func (ctx *requestCtx) Value(key any) any                       { return nil }

func (ctx *requestCtx) Err() error {
	select {
	case <-ctx.Done():
		return context.Canceled
	default:
		return nil
	}
}

var ctx *requestCtx

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	buf := &bytes.Reader{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
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
