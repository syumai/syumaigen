package syumaigen

import "image/color"

type ColorMap map[int]color.Color

var DefaultColorMap = map[int]color.Color{
	0: color.Transparent,
	1: color.Black,
	2: color.Black,
	3: color.White,
	4: color.White,
	5: color.RGBA{255, 0, 0, 255},
	6: color.RGBA{0, 255, 0, 255},
	7: color.RGBA{0, 255, 0, 255},
	8: color.RGBA{0, 255, 0, 255},
	9: color.RGBA{0, 255, 0, 255},
}
