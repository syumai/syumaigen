package syumaigen

import (
	"bytes"
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
)

func GenerateSVG(data [][]int, cmap ColorMap, scale int) (io.Reader, error) {
	if err := assertData(data); err != nil {
		return nil, err
	}
	if scale < 1 {
		return nil, fmt.Errorf("scale must be >= 1")
	}
	width := len(data[0]) * scale
	height := len(data) * scale

	var buf bytes.Buffer
	canvas := svg.New(&buf)
	canvas.Start(width, height)

	for i, line := range data {
		for j, n := range line {
			c, ok := cmap[n]
			if !ok {
				return nil, fmt.Errorf("color not found: %d", n)
			}
			cc, ok := colorful.MakeColor(c)
			if !ok {
				continue
			}
			canvas.Square(j*scale, i*scale, scale, fmt.Sprintf("fill: %s", cc.Hex()))
		}
	}
	canvas.End()
	return &buf, nil
}
