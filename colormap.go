package syumaigen

import (
	"image/color"
	"math/rand"
	"strings"
	"time"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type ColorMap map[int]color.Color

var DefaultColorMap = ColorMap{
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

func GenerateRandomColorMap() ColorMap {
	rand.Seed(time.Now().UnixNano())
	h := rand.Float64() * 360.0
	c := 0.4 + rand.Float64()*0.6
	return GenerateColorMapByHCL(h, c)
}

func GenerateColorMapByHCL(h float64, c float64) ColorMap {
	return ColorMap{
		0: DefaultColorMap[0],
		1: DefaultColorMap[1],
		2: DefaultColorMap[2],
		3: DefaultColorMap[3],
		4: DefaultColorMap[4],
		5: DefaultColorMap[5],
		6: colorful.Hcl(h, c, 0.5).Clamped(),
		7: colorful.Hcl(h, c, 0.3).Clamped(),
		8: colorful.Hcl(h, c, 0.7).Clamped(),
		9: colorful.Hcl(h, c, 0.9).Clamped(),
	}
}

func GenerateColorMapByColorCode(cc string) ColorMap {
	if !strings.HasPrefix(cc, "#") {
		cc = "#" + cc
	}
	col, err := colorful.Hex(cc)
	if err != nil {
		return DefaultColorMap
	}
	h, c, l := col.Hcl()
	return ColorMap{
		0: DefaultColorMap[0],
		1: DefaultColorMap[1],
		2: DefaultColorMap[2],
		3: DefaultColorMap[3],
		4: DefaultColorMap[4],
		5: DefaultColorMap[5],
		6: colorful.Hcl(h, c, l).Clamped(),
		7: colorful.Hcl(h, c, l*0.8).Clamped(),
		8: colorful.Hcl(h, c, 1.0-(1.0-l)*0.8).Clamped(),
		9: colorful.Hcl(h, c, 1.0-(1.0-l)*0.2).Clamped(),
	}
}
