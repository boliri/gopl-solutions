// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
// Elements can be selected by either specifying:
//   - Their name, or
//   - Their name and one or more attributes
//
// Examples:
// - Selecting by name:
//   - div
//   - div div h2
//
// - Selecting by name and attributes:
//   - div:id=foo
//   - div:id=foo,class=bar
//   - div:id=foo,class=bar div h2
//   - div div:id=foo,class=bar h2
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type selector xml.StartElement

func main() {
	dec := xml.NewDecoder(os.Stdin)

	in := os.Args[1:]
	selectors := make([]selector, len(in))
	err := parseSelectors(&selectors, in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}

	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, selectors) {
				fmt.Printf("%s: %s\n", prettify(stack), tok)
			}
		}
	}
}

// parseSelectors loads to s all selectors from raw as Selector instances
// raw's elements must comply with format "name:attrname=attrval,attrname=attrval"
func parseSelectors(s *[]selector, raw []string) error {
	var parsed []string

	for i, rsel := range raw {
		parsed = strings.Split(rsel, ":")

		name := parsed[0]
		sel := selector{
			Name: xml.Name{Space: "", Local: name},
			Attr: []xml.Attr{},
		}

		// name based selector
		if len(parsed) == 1 {
			(*s)[i] = sel
			continue
		}

		// name & attrs based selector
		if parsed[1] == "" {
			return fmt.Errorf("parse: malformed selector \"%s\": no attributes provided", rsel)
		}

		rattrs := strings.Split(parsed[1], ",")
		for _, ra := range rattrs {
			if ra == "" {
				return fmt.Errorf("parse: malformed selector \"%s\": one or more attribute selectors are missing", rsel)
			}

			parsed = strings.Split(ra, "=")
			if parsed[0] == "" {
				return fmt.Errorf("parse: malformed attr selector \"%s\": missing name", ra)
			}

			if parsed[1] == "" {
				return fmt.Errorf("parse: malformed attr selector \"%s\": missing value", ra)
			}

			k, v := parsed[0], parsed[1]
			sel.Attr = append(sel.Attr, xml.Attr{
				Name:  xml.Name{Space: "", Local: k},
				Value: v,
			})
		}

		(*s)[i] = sel
	}

	return nil
}

// prettify returns a string with a pretty representation of multiple StartElement instances
func prettify(e []xml.StartElement) string {
	var b strings.Builder
	for i, elem := range e {
		b.WriteString(elem.Name.Local)

		if len(elem.Attr) != 0 {
			b.WriteString(":")
			for j, attr := range elem.Attr {
				b.WriteString(fmt.Sprintf("%s=%s", attr.Name.Local, attr.Value))
				if j != len(elem.Attr)-1 {
					b.WriteString(",")
				}
			}
		}

		if i != len(e)-1 {
			b.WriteString(" ")
		}
	}

	return b.String()
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []selector) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0].Name == y[0].Name {
			found, expected := 0, len(y[0].Attr)
			for _, selAttr := range y[0].Attr {
				for _, attr := range x[0].Attr {
					if selAttr.Name.Local == attr.Name.Local && selAttr.Value == attr.Value {
						found += 1
					}
				}
			}
			if found != expected {
				return false
			}
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
