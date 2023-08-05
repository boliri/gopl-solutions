package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo1 version that prints the executed command along with its command-line arguments.
func echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Echo2 version that prints the executed command along with its command-line arguments.
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Echo3 version that prints the executed command along with its command-line arguments.
func echo3() {
	fmt.Println(strings.Join(os.Args, " "))
}

func main() {
	echo1()
	echo2()
	echo3()
}
