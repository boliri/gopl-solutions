// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"log"
	"os"

	"wordsimages"
)

func main() {
	url := os.Args[1]
	words, images, err := wordsimages.CountWordsAndImages(url)
	if err != nil {
		log.Fatal("an error occurred while counting words and images in %s: %s", url, err)
	}

	fmt.Printf("Found %d words and %d images in %s HTML document\n\n", words, images, url)
}

//!-
