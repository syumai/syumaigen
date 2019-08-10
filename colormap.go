package syumaigen

import "image/color"

type ColorMap map[int]color.Color

var DefaultColorMap = map[int]color.Color{
	0: color.Transparent,
	1: color.Black,
	2: color.RGBA{66, 66, 66, 255},
	3: color.RGBA{255, 255, 240, 255},
	4: color.RGBA{222, 222, 203, 255},
	5: color.RGBA{255, 121, 0, 255},
	6: color.RGBA{28, 214, 1, 255},
	7: color.RGBA{25, 179, 3, 255},
	8: color.RGBA{126, 214, 113, 255},
	9: color.RGBA{191, 214, 188, 255},
}
