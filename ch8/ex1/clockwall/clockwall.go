package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type clock struct {
	location, server, time string
}

func (c *clock) sync() {
	conn, err := net.Dial("tcp", c.server)
	if err != nil {
		log.Fatalf("cannot sync %s's time: %s", c.location, err)
	}
	defer conn.Close()

	if _, err := io.Copy(c, conn); err != nil {
		log.Fatal(err)
	}
}

func (c *clock) Write(p []byte) (n int, err error) {
	c.time = string(p)
	return len(p), nil
}

func main() {
	var clocks []*clock

	err := parseArgs(&clocks)
	if err != nil {
		log.Fatal(err)
	}

	startSyncing(clocks)
	waitForClockServers(clocks)
	printTimes(clocks)
}

func parseArgs(cs *[]*clock) error {
	for _, arg := range os.Args[1:] {
		fragments := strings.Split(arg, "=")
		if len(fragments) != 2 {
			return fmt.Errorf("error: parse: invalid format: \"%s\"", arg)
		}

		(*cs) = append((*cs), &clock{location: fragments[0], server: fragments[1]})
	}

	return nil
}

func startSyncing(cs []*clock) {
	for _, c := range cs {
		go c.sync()
	}
}

func waitForClockServers(cs []*clock) {
	desynced := make([]*clock, len(cs))
	copy(desynced, cs)
	for len(desynced) > 0 {
		c := desynced[0]
		if c.time != "" {
			desynced = desynced[1:]
		}
	}
}

func printTimes(cs []*clock) {
	tabw := &tabwriter.Writer{}
	tabw.Init(os.Stdout, 8, 8, 0, '\t', 0)

	for {
		fmt.Fprintf(tabw, "\n %s\t%s\t", "Location", "Time")
		fmt.Fprintf(tabw, "\n %s\t%s\t", "--------", "----")
		for _, c := range cs {
			fmt.Fprintf(tabw, "\n %s\t%s\t", c.location, c.time)
		}
		tabw.Flush()

		time.Sleep(1 * time.Second)
	}
}
