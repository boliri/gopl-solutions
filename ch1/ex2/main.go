package main

import (
	"fmt"
	"os"
)

// Echo1 version that prints all the command-line arguments and their position.
func echo1() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d %s\n", i, os.Args[i])
	}
}

// Echo2 version that prints all the command-line arguments and their position.
func echo2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i + 1, arg)
	}
}

func main() {
	echo1()
	echo2()
}
