package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"

	"reader"
)

func main() {
	s := `
	<html>
	    <head>
		    <title>Fancy HTML</title>
		</head>
		<body>
			<div>Hey!</div>
		</body>
	</html>
	`

	r := reader.NewReader(s)

	doc, _ := html.Parse(r) // parse errors skipped for brevity

	fmt.Print("HTML parsed successfully. Let's render it:\n\n")
	html.Render(os.Stdout, doc)
}
