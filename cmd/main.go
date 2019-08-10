package main

import (
	"image/png"
	"log"
	"os"

	"github.com/syumai/syumaigen"
)

func main() {
	img, err := syumaigen.GenerateImage(
		syumaigen.Pattern,
		syumaigen.DefaultColorMap,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(os.Stdout, img); err != nil {
		log.Fatal(err)
	}
}
