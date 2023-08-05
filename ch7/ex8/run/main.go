package main

import (
	"fmt"
	"os"
	"sort"

	"multitiersort"
)

func main() {
	tracks := []*multitiersort.Track{
		multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2013, "3m39s"),
		multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m39s"),
		multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
	}
	ct := multitiersort.NewClickableTable(tracks)

	fmt.Print("Original table:\n\n")
	ct.Print(os.Stdout)
	fmt.Printf("\n--------\n\n")

	fmt.Print("Sorting table by length...\n\n")
	ct.ClickColumnHead("length")
	sort.Sort(ct.Mts)
	ct.Print(os.Stdout)
	fmt.Printf("\n--------\n\n")

	fmt.Print("Sorting table by length and year...\n\n")
	ct.ClickColumnHead("year")
	sort.Sort(ct.Mts)
	ct.Print(os.Stdout)
}
