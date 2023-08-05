package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"mandelbrot"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/mandelbrot", mandelbrotHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// plots a mandelbrot
func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
		return
	}

	x, err := strconv.Atoi(r.Form["x"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	y, err := strconv.Atoi(r.Form["y"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	zoom, err := strconv.Atoi(r.Form["zoom"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	cfg := mandelbrot.NewMandelbrotConfig(x, y, zoom)

	mandelbrot.GeneratePNG(w, cfg)
}