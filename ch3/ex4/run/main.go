package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"surface"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/surface", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// plots a surface
func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
		return
	}

	width, err := strconv.Atoi(r.Form["width"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	height, err := strconv.Atoi(r.Form["height"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	cells, err := strconv.Atoi(r.Form["cells"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	fn := strings.Join(r.Form["fn"], "")

	var f surface.PlotFn
	switch fn {
		case "eggbox":
			f = surface.Eggbox
		case "saddle":
			f = surface.Saddle
		case "mogguls":
			f = surface.Mogguls
		default:
			f = surface.F
	}

	cfg := surface.NewSurfaceConfig(width, height, cells, f)

	w.Header().Set("Content-Type", "image/svg+xml")
	surface.GenerateSVG(w, cfg)
}