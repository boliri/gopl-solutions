package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"xmltree"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)

	var root xmltree.Node
	var stack []*xmltree.Element

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmltree: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := xmltree.Element{
				Name:     tok.Name,
				Attr:     tok.Attr,
				Children: []xmltree.Node{},
			}

			if root == nil {
				root = &e
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, &e)
			}

			stack = append(stack, &e)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			// sometimes decoders can yield  space-like chars right before reaching
			// EOF - these should be ignored to prevent runtime errors when resizing the
			// stack slice
			if len(stack) == 0 {
				continue
			}

			e := xmltree.CharData(string(tok))
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, &e)
		}
	}

	fmt.Println(root)
}

//!-
