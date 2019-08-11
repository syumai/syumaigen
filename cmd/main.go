package main

import (
	"flag"
	"image/gif"
	"image/png"
	"log"
	"os"

	"github.com/syumai/syumaigen"
)

var (
	scale    = flag.Int("scale", 10, "specify image scale")
	random   = flag.Bool("random", true, "randomize color generation")
	animated = flag.Bool("animated", false, "generate animated GIF")
)

func main() {
	flag.Parse()

	if *animated {
		img, err := syumaigen.GenerateAnimatedSyumaiGIF(*scale)
		if err != nil {
			log.Fatal(err)
		}
		if err := gif.EncodeAll(os.Stdout, img); err != nil {
			log.Fatal(err)
		}
		return
	}

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
