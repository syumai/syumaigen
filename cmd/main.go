package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/syumai/syumaigen"
)

var (
	scale  = flag.Int("scale", 10, "specify image scale")
	random = flag.Bool("random", true, "randomize color generation")
)

func main() {
	flag.Parse()

	colorMap := syumaigen.DefaultColorMap
	if *random {
		colorMap = syumaigen.GenerateRandomColorMap()
	}

	img, err := syumaigen.GenerateImage(
		syumaigen.Pattern,
		colorMap,
		*scale,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}
