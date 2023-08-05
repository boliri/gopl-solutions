// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var portFlag *int = flag.Int("port", 8000, "port to listen for incoming connections")
var tzFlag *string = flag.String("tz", "UTC", "timezone for which current time should be yielded for")

func handleConn(c net.Conn) {
	defer c.Close()

	tz, err := time.LoadLocation(*tzFlag)
	if err != nil {
		panic(fmt.Sprintf("timezone \"%s\" is not valid", tz))
	}

	for {
		_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portFlag))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
