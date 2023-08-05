package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"bitsdiffsha256"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]

	h1, h2 := sha256.Sum256([]byte(s1)), sha256.Sum256([]byte(s2))
	diff := bitsdiffsha256.GetBitsDiff(h1, h2)

	hex1, hex2 := fmt.Sprintf("%x", h1), fmt.Sprintf("%x", h2)

	fmt.Println("Input #1")
	fmt.Printf("\tString: %s\n", s1)
	fmt.Printf("\tSHA-256 hash: %s\n\n", hex1)

	fmt.Println("Input #2")
	fmt.Printf("\tString: %s\n", s2)
	fmt.Printf("\tSHA-256 hash: %s\n\n", hex2)

	fmt.Printf("Both hashes differ in %d bits\n\n", diff)
}
