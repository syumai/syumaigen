package main

import (
	"flag"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/syumai/syumaigen"
)

var (
	scale     = flag.Int("scale", 10, "specify image scale")
	code      = flag.String("code", "", "use color code")
	norandom    = flag.Bool("norandom", false, "stop randomize color generation")
	animated  = flag.Bool("animated", false, "generate animated GIF")
	outputSVG = flag.Bool("svg", false, "generate SVG")
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
	if *code != "" {
		colorMap = syumaigen.GenerateColorMapByColorCode(*code)
	} else if !*norandom {
		colorMap = syumaigen.GenerateRandomColorMap()
	}

	if *outputSVG {
		img, err := syumaigen.GenerateSVG(
			syumaigen.Pattern,
			colorMap,
			*scale,
		)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.Copy(os.Stdout, img); err != nil {
			log.Fatal(err)
		}
		return
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
