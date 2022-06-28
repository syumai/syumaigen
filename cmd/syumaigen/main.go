package main

import (
	"flag"
	"image/color"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/syumai/syumaigen"
)

var (
	scale     = flag.Int("scale", 10, "specify image scale")
	code      = flag.String("code", "", "use color code")
	bgCode    = flag.String("bgcode", "", "use background color code")
	norandom  = flag.Bool("norandom", false, "stop randomize color generation")
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

	if *bgCode != "" {
		c, err := parseColorCode(*bgCode)
		if err != nil {
			log.Fatal(err)
		}
		colorMap[0] = c
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

func parseColorCode(cc string) (color.Color, error) {
	if !strings.HasPrefix(cc, "#") {
		cc = "#" + cc
	}
	col, err := colorful.Hex(cc)
	if err != nil {
		return nil, err
	}
	h, c, l := col.Hcl()
	return colorful.Hcl(h, c, l).Clamped(), nil
}
