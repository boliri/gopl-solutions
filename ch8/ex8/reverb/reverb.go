// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const tickrate = time.Duration(10 * time.Second)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(s *bufio.Scanner, ch chan<- bool) {
	for s.Scan() {
		ch <- true
	}
}

// !+
func handleConn(c net.Conn) {
	defer c.Close() // NOTE: ignoring potential errors from input.Err()

	input := bufio.NewScanner(c)
	scanCh := make(chan bool, 1)
	go scan(input, scanCh)

	ticker := time.NewTicker(tickrate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			return
		case <-scanCh:
			go echo(c, input.Text(), 1*time.Second)
			ticker.Reset(tickrate)
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
