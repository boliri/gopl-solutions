// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import "fmt"

func run(arg interface{}) (res interface{}) {
	defer func() { res = recover() }()
	panic(arg)
}


func main() {
	fmt.Println(run("foo"))
}
