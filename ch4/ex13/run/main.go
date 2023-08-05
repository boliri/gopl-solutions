// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

package main

import (
	"fmt"
	"os"

	"poster"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		if args[i] == "-h" || args[i] == "--help" {
			help()
			return
		}
	}

	pArgs := parseArgs(args)
	req := &poster.PosterRequest{
		ApiKey: pArgs["api-key"],
		Title: pArgs["movie"],
	}

	err := poster.Download(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Poster image successfully downloaded for movie %s\n\n", req.Title)
}

func help() {
	fmt.Println("This tool lets you download image posters of movies by its name from the OMDB API.")
	fmt.Println("\t-h | --help: prints this help message")
	fmt.Println("\t-k | --api-key: key needed to make requests to this API, you can grab one from their website for free")
	fmt.Println("\t-m | --movie: title of the movie you want to download a poster for")
}

func parseArgs(args []string) map[string]string {
	result := make(map[string]string)

	for i := 0; i < len(args); {
		if args[i] == "-k" || args[i] == "--api-key" {
			result["api-key"] = args[i+1]
		} else if args[i] == "-m" || args[i] == "--movie" {
			result["movie"] = args[i+1]
		}
		i += 2
	}

	return result
}

//!-
