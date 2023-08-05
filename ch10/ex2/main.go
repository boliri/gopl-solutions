// This program reads zip and POSIX tar files from the command line and prints their contents
package main

import (
	"arch"
	_ "arch/tar"
	_ "arch/zip"
	"fmt"
	"os"
)

func main() {
	a, kind, err := arch.Read(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(kind, "file:")
	fmt.Println(a)
}
