// Extend the jpeg program so that it converts any supported input format to any output format,
// using image.Decode to detect the input format and a flag to select the output format.
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
	"strings"
)

var outputFmt string

func init() {
	flag.StringVar(&outputFmt, "fmt", "jpeg", "output image format")
	flag.Parse()
}

func main() {
	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch strings.ToLower(outputFmt) {
	case "jpeg", "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	default:
		return fmt.Errorf("unknown format %q", outputFmt)
	}
}
