package wordsimages

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	visit(&words, &images, n)
	return
}

// visit explores a HTML document tree and keeps track of the number of words and images in it
func visit(words, images *int, n *html.Node) {
	if n.Type == html.TextNode && !unicode.IsSpace(rune(n.Data[0])) {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			*words++
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(words, images, c)
	}
}