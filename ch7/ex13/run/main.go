package main

import (
	"fmt"

	"eval"
)

func main() {
	var e string
	var ex eval.Expr

	e = "x + 1"
	ex, _ = eval.Parse(e) // skipping error checks for brevity
	fmt.Printf("Original expression: %s\n", e)
	fmt.Printf("Pretty-printed expression: %s\n\n", ex)

	e = "x - y"
	ex, _ = eval.Parse(e) // skipping error checks for brevity
	fmt.Printf("Original expression: %s\n", e)
	fmt.Printf("Pretty-printed expression: %s\n\n", ex)

	e = "sqrt(x * 2 + 1)"
	ex, _ = eval.Parse(e) // skipping error checks for brevity
	fmt.Printf("Original expression: %s\n", e)
	fmt.Printf("Pretty-printed expression: %s\n\n", ex)
}
