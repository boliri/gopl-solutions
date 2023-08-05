package prettyhtml

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettifiedHtmlCanBeParsed(t *testing.T) {
    in := "<html><head></head><body><div id='a'><div id='b' class='boom'>hey!</div></div><!--this is a comment--><img source='foo'></img></body></html>"
	inreader := strings.NewReader(in)
	indoc, _ := html.Parse(inreader)

	out := Prettify(indoc)
	outreader := strings.NewReader(out)
	_, err := html.Parse(outreader)

    if err != nil {
        t.Fatalf(`could not parse prettified html string to html.Node`)
    }
}