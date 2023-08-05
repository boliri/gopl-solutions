package multitiersort

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"golang.org/x/exp/slices"
)

// Track

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func NewTrack(title string, artist string, album string, year int, trackLen string) *Track {
	return &Track{title, artist, album, year, length(trackLen)}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// Sorting

type multiTierSort struct {
	table    []*Track
	criteria []string
}

func (mts multiTierSort) Len() int {
	return len(mts.table)
}

func (mts multiTierSort) Less(i, j int) bool {
	return mts.lessByCriteria(mts.table[i], mts.table[j])
}

func (mts multiTierSort) Swap(i, j int) {
	mts.table[i], mts.table[j] = mts.table[j], mts.table[i]
}

func (mts *multiTierSort) lessByCriteria(x, y *Track) bool {
	for _, c := range mts.criteria {
		if c == "title" && x.Title != y.Title {
			return x.Title < y.Title
		}

		if c == "artist" && x.Artist != y.Artist {
			return x.Artist < y.Artist
		}

		if c == "album" && x.Album != y.Album {
			return x.Album < y.Album
		}

		if c == "year" && x.Year != y.Year {
			return x.Year < y.Year
		}

		if c == "length" && x.Length != y.Length {
			return x.Length < y.Length
		}
	}
	return false
}

// Clickable table
type clickableTable struct {
	Mts multiTierSort
}

func NewClickableTable(table []*Track) *clickableTable {
	return &clickableTable{multiTierSort{table: table}}
}

func (t *clickableTable) ClickColumnHead(id string) {
	// Do nothing if the criteria is actively used already
	if slices.Contains(t.Mts.criteria, id) {
		return
	}

	t.Mts.criteria = append(t.Mts.criteria, id)
}

func (t *clickableTable) Print(w io.Writer) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(w, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range t.Mts.table {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate columns widths and print table
}
