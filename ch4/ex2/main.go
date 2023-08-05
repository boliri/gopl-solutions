package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]

	var hash string = fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
	if len(os.Args) > 2 && os.Args[2] == "--hash" {
		if os.Args[3] == "sha384" {
			hash = fmt.Sprintf("%x", sha512.Sum384([]byte(input)))
		} else if os.Args[3] == "sha512" {
			hash = fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
		}
	}

	fmt.Println(hash)
}
