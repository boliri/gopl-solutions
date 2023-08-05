package xmltree

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Node interface{}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Name     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	var b strings.Builder

	// opening tag
	b.WriteString(fmt.Sprintf("<%s", e.Name.Local))
	for _, a := range e.Attr {
		b.WriteString(fmt.Sprintf(" %s=\"%s\"", a.Name.Local, a.Value))
	}
	b.WriteString(">")

	// body
	for _, child := range e.Children {
		switch child := child.(type) {
		case *Element:
			b.WriteString(child.String())
		case *CharData:
			b.WriteString(child.String())
		}
	}

	// closing tag
	b.WriteString(fmt.Sprintf("</%s>", e.Name.Local))

	return b.String()
}
