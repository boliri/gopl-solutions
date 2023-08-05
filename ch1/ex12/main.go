// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type LissajousConfig struct {
	Cycles			int
	Resolution  	float32
	Size			int
	NumberOfFrames  int
	Delay			int
}

func newDefaultLissajousConfig() LissajousConfig {
	return LissajousConfig{
		Cycles: 5,
		Resolution: 0.001,
		Size: 100,
		NumberOfFrames: 64,
		Delay: 8,
	}
}

func newLissajousConfig(cycles int, numberOfFrames int) LissajousConfig {
	cfg := newDefaultLissajousConfig()
	cfg.Cycles = cycles
	cfg.NumberOfFrames = numberOfFrames

	return cfg
}


var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// lissajous builds a GIF based on optional query params passed along with the URL.
func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
		return
	}

	cycles, err := strconv.Atoi(r.Form["cycles"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	frames, err := strconv.Atoi(r.Form["frames"][0])
	if err != nil {
		fmt.Fprint(w, "Error while parsing query params: ", err)
	}

	cfg := newLissajousConfig(cycles, frames)
	lissajous(w, cfg)
}

func lissajous(out io.Writer, cfg LissajousConfig) {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: cfg.NumberOfFrames}
	phase := 0.0 // phase difference
	for i := 0; i < cfg.NumberOfFrames; i++ {
		rect := image.Rect(0, 0, 2*cfg.Size+1, 2*cfg.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cfg.Cycles)*2*math.Pi; t += float64(cfg.Resolution) {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(cfg.Size+int(x*float64(cfg.Size)+0.5), cfg.Size+int(y*float64(cfg.Size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, cfg.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-