package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"textnodes"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textnodes: %v\n", err)
		os.Exit(1)
	}

	var texts []string
	textnodes.Find(&texts, doc)

	for _, t := range texts {
		fmt.Println(t)
	}
}
