package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

// enables error printing
var verbose *bool = flag.Bool("v", false, "verbose mode")

// dirChans maps filesystem dirs to buffered channels with capacity = 1, where each channel
// holds the current disk usage of that dir materialized as a diskUsage
type dirChans map[string]chan *diskUsage

// closeAll closes all channels in d
func (d dirChans) closeAll() {
	for _, ch := range d {
		close(ch)
	}
}

// drainAll drains all channels in d
func (d dirChans) drainAll() {
	for _, ch := range d {
		for range ch {
			// do nothing
		}
	}
}

// flattenDiskUsages drains all channels in d and returns a new diskUsage with
// the total disk usage of all its dirs
func (d dirChans) flattenDiskUsages() diskUsage {
	totals := diskUsage{}
	for _, ch := range d {
		for du := range ch {
			totals.nfiles += du.nfiles
			totals.nbytes += du.nbytes
		}
	}
	return totals
}

// diskUsage holds files and bytes counters
type diskUsage struct {
	nfiles, nbytes int64
}

// !+1
var done = make(chan struct{})
var cancel = make(chan struct{})

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

//!-1

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// initialize root dirs to chans mapping
	dc := dirChans{}

	//!+2
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()
	//!-2

	// Traverse each root of the file tree in parallel.
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)

		// populate dc with current root dir and a buffered channel
		fileSizes := make(chan *diskUsage, 1)
		dc[root] = fileSizes

		// push a zero diskUsage to the root dir's channel
		// this instance will be updated on every walked dir if a file is found
		fileSizes <- &diskUsage{}

		// now, let's walk the root dir
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		done <- struct{}{}
	}()

	// Print the results periodically.
	tick := time.Tick(500 * time.Millisecond)

	// let's use a tabwriter to print results on every tick nicely
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)

	//!+3
	for {
		select {
		case <-done:
			dc.closeAll()

			totalDu := dc.flattenDiskUsages()
			printDiskUsage(totalDu.nfiles, totalDu.nbytes)
			fmt.Println()

			return
		case <-cancel:
			dc.closeAll()
			dc.drainAll()
			return
		case <-tick:
			// print current disk usage per root dir
			// note that the root dir's usage is pushed to that dir's channel immediatelly after
			// polling it, so walkDir goroutines can keep going
			usages := make([]string, 0, len(dc))
			for dir, c := range dc {
				usage := <-c
				c <- usage
				usages = append(usages, getDiskUsageForDir(dir, usage.nfiles, usage.nbytes))
			}
			sort.Strings(usages) // ensure order consistency across successive ticks
			w.Write([]byte(strings.Join(usages, "\n")))
			w.Write([]byte("\n\n"))
			w.Flush()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB", nfiles, float64(nbytes)/1e9)
}

func getDiskUsageForDir(dir string, nfiles, nbytes int64) string {
	return fmt.Sprintf("%s\t  %d files\t  %.1f GB", dir, nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and updates the diskUsage instance in fileSizes if a file is found.
// !+4
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan *diskUsage) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		// ...
		//!-4
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			// read from fileSizes, update its diskUsage, and push it again to fileSizes
			du := <-fileSizes
			du.nfiles++
			du.nbytes += entry.Size()
			fileSizes <- du
		}
		//!+4
	}
}

//!-4

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
// !+5
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-cancel:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	// ...read directory...
	//!-5

	f, err := os.Open(dir)
	if err != nil && *verbose {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil && *verbose {
		// fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
