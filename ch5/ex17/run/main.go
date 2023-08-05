package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"

	"findelems"
)

func main() {
	in, lookup := os.Args[1], os.Args[2:]

	if len(lookup) == 0 {
		log.Fatal("no tag names provided")
	}

	inreader := strings.NewReader(in)
	doc, err := html.Parse(inreader)
	if err != nil {
		log.Fatal(err)
	}

	found := findelems.ElementsByTagName(doc, lookup...)
	if len(found) == 0 {
		fmt.Printf("no elements match any of %v\n", lookup)
	} else {
		fmt.Printf("%d element(s) found:\n\n", len(found))
		for _, elem := range found {
			parts := []string{elem.Data}
			for _, a := range elem.Attr {
				parts = append(parts, fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
			}

			fmt.Printf("<%s>...</%s>\n", strings.Join(parts, " "), elem.Data)
		}
	}
}