package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	"links"

	"golang.org/x/net/html"
)

var domain string

// !+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 5)

// getMirror returns a mirrorized version of the root domain
func getMirror() string {
	return "mirror." + domain
}

// get makes a HTTP GET request to the specified URL, and returns the server response as is
// the response is NOT closed, so its body can be consumed from the caller
func get(url string) (*http.Response, error) {
	tokens <- struct{}{} // acquire a token
	resp, err := http.Get(url)
	<-tokens // release the token

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	return resp, err
}

// crawl looks for links in a HTML document, and after extracting them, they're mutated to point to the
// mirror website in the HTML document
func crawl(list *[]string, doc *html.Node, url *url.URL) {
	preFn := links.ExtractLinksWrapper(list, url)

	transformFn := func(hrefVal string) string { return strings.ReplaceAll(hrefVal, url.Hostname(), getMirror()) }
	postFn := links.TransformHrefsWrapper(transformFn, url)

	links.ProcessDocument(doc, preFn, postFn)
}

//!-sema

// !+
func main() {
	flag.Parse()

	worklist := make(chan string)
	var wg sync.WaitGroup

	raw := os.Args[1]
	url, err := url.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}

	domain = url.Hostname()

	// create base folder in local disk
	err = os.Mkdir(getMirror(), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// we need to increment the waitgroup so the worklist channel is not closed prematurely in one of the upcoming goroutines
	wg.Add(1)

	go func() { worklist <- raw }()

	// Close the worklist channel when there are no more urls to fetch
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for url := range worklist {
		if !seen[url] {
			seen[url] = true

			// do not increment the waitgroup if the url is the initial one
			// otherwise the worklist channel will never get closed and we'll fall into a deadlock
			if url != raw {
				wg.Add(1)
			}

			go func(url string) {
				defer wg.Done()

				// make HTTP GET request
				resp, err := get(url)
				if err != nil {
					fmt.Printf("%s: %s\n", url, err)
					return
				}
				defer resp.Body.Close()

				// write response body to disk right away if not an HTML document
				if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
					filepath := path.Join(getMirror(), resp.Request.URL.Path)
					bytes, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Printf("reading response from %s: %s", resp.Request.URL, err)
						return
					}
					os.WriteFile(filepath, bytes, os.ModePerm)
					return
				}

				// try to parse the response body as HTML
				doc, err := html.Parse(resp.Body)
				if err != nil {
					fmt.Printf("parsing %s as HTML: %v\n", url, err)
					return
				}

				// crawl the web and mutate links found in the HTML document so they point to
				// the mirror website
				var linksList []string
				crawl(&linksList, doc, resp.Request.URL)

				// write the response body to disk
				var filepath string
				if resp.Request.URL.Path == "" || resp.Request.URL.Path == "/" {
					filepath = path.Join(getMirror(), "index.html")
				} else {
					filepath = path.Join(getMirror(), resp.Request.URL.Path)
					if !strings.HasSuffix(filepath, ".html") {
						filepath += ".html"
					}
				}

				var w bytes.Buffer
				html.Render(&w, doc)
				bytes, err := ioutil.ReadAll(&w)
				if err != nil {
					fmt.Printf("rendering HTML for %s: %s", url, err)
					return
				}
				os.WriteFile(filepath, bytes, os.ModePerm)

				// push new urls to the worklist, but only if they are part of the original domain
				// CAVEAT: subdomains are also crawled! a deep refactor to work with url.URL instances
				// instead of plain strings should mitigate this - however, and technically speaking,
				// subdomains are part of a domain...aren't they?
				for _, u := range linksList {
					if !strings.Contains(u, domain) {
						continue
					}
					worklist <- u
				}
			}(url)
		}
	}
}

//!-
