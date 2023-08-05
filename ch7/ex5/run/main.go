package main

import (
	"fmt"
	"io"
	"strings"

	"reader"
)

func main() {
	s := "Hi there! How are you?"
	strR := strings.NewReader(s)

	var limit int64 = 10
	r := reader.LimitReader(strR, limit)

	b := make([]byte, len(s))
	n, err := r.Read(b)

	fmt.Printf("%d bytes read from buffer. String is: \"%s\"\n", n, string(b))
	if err == io.EOF {
		fmt.Println("EOF reached")
	}
}
