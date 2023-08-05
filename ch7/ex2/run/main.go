package main

import (
	"bytes"
	"fmt"

	"counter"
)

func main() {
	w := bytes.NewBuffer(nil)
	bc, written := counter.CountingWriter(w)

	s := "Hi there! How are you?"
	fmt.Printf("\"%s\"\n\n", s)

	bc.Write([]byte(s))
	fmt.Printf("%d bytes written to buffer %T so far\n\n", *written, w)
}
