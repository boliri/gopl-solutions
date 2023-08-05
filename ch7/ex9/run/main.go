package main

import (
	"log"
	"net/http"
	"sort"

	"multitiersort"
)

var tracks = []*multitiersort.Track{
	multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2013, "3m39s"),
	multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m39s"),
	multitiersort.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
}
var clickableTable = multitiersort.NewClickableTable(tracks)

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/sort", sortByColumn)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func landing(w http.ResponseWriter, r *http.Request) {
	clickableTable.PrintHtml(w)
}

func sortByColumn(w http.ResponseWriter, r *http.Request) {
	clickableTable.ClickColumnHead(r.URL.Query().Get("column"))
	sort.Sort(clickableTable.Mts)
	clickableTable.PrintHtml(w)
}
