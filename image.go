package syumaigen

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
)

var transparentPalette []color.Color

func init() {
	transparentPalette = append(transparentPalette, palette.WebSafe...)
	transparentPalette = append(transparentPalette, color.RGBA{0, 0, 0, 0})
}

func assertData(data [][]int) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("data is blank")
	}
	lineLen := len(data[0])
	for _, line := range data {
		if lineLen != len(line) {
			return fmt.Errorf("line length is not equal, want: %d, got: %d", lineLen, len(line))
		}
	}
	return nil
}

func GenerateImage(data [][]int, cmap ColorMap, scale int) (image.Image, error) {
	if err := assertData(data); err != nil {
		return nil, err
	}
	if scale < 1 {
		return nil, fmt.Errorf("scale must be >= 1")
	}
	img := image.NewRGBA(image.Rect(0, 0, len(data[0])*scale, len(data)*scale))
	for i, line := range data {
		for j, n := range line {
			c, ok := cmap[n]
			if !ok {
				return nil, fmt.Errorf("color not found: %d", n)
			}
			for xs := 0; xs < scale; xs++ {
				for ys := 0; ys < scale; ys++ {
					img.Set(j*scale+ys, i*scale+xs, c)
				}
			}
		}
	}
	return img, nil
}

func GenerateAnimatedSyumaiGIF(scale int) (*gif.GIF, error) {
	g := &gif.GIF{}
	frames := 30
	for i := 0; i < frames; i++ {
		h := float64(i) / float64(frames) * 360.0
		img, err := GenerateImage(
			Pattern,
			GenerateColorMapByHCL(h, 0.95),
			scale,
		)
		if err != nil {
			return nil, err
		}
		palettedImg := image.NewPaletted(img.Bounds(), transparentPalette)
		draw.FloydSteinberg.Draw(palettedImg, img.Bounds(), img, image.ZP)
		g.Image = append(g.Image, palettedImg)
		g.Delay = append(g.Delay, 10)
	}
	return g, nil
}
