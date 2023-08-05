// The imgconv command reads an arbitrary image from the standard input
// and writes it as an image with the desired output format.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outFmt = flag.String("out-fmt", "png", "output format")

var supportedFmts = map[string]bool{
	"jpeg": true,
	"gif":  true,
	"png":  true,
}

func main() {
	flag.Parse()

	if !supportedFmts[*outFmt] {
		fmt.Fprintln(os.Stderr, "imgconv: unsupported output format")
		os.Exit(1)
	}

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "imgconv:", err)
		os.Exit(1)
	}

	if kind == *outFmt {
		fmt.Fprintf(os.Stderr, "imgconv: image is already a %s\n", *outFmt)
		os.Exit(0)
	}

	fmt.Fprintln(os.Stderr, "Input format =", kind)
	fmt.Fprintln(os.Stderr, "Output format =", *outFmt)
	if err := toOutFmt(os.Stdout, img); err != nil {
		fmt.Fprintf(os.Stderr, "imgconv: %s -> %s: %v\n", kind, *outFmt, err)
		os.Exit(1)
	}
}

func toOutFmt(out io.Writer, in image.Image) error {
	switch *outFmt {
	case "png":
		return png.Encode(out, in)
	case "gif":
		return gif.Encode(out, in, nil)
	case "jpeg":
		return jpeg.Encode(out, in, nil)
	default:
		panic(fmt.Sprintf("imgconv: unexpected output format %s", *outFmt))
	}
}
