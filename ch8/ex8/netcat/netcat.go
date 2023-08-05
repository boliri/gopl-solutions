// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

// !+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("connection was closed due to inactivity")
		done <- struct{}{} // signal the main goroutine
	}()

	go mustCopy(conn, os.Stdin, done)
	<-done
}

//!-

func mustCopy(dst io.Writer, src io.Reader, done chan struct{}) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
		done <- struct{}{}
	}
}
