package syumaigen

import (
	"fmt"
	"image"
)

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

func GenerateImage(data [][]int, cmap ColorMap) (image.Image, error) {
	if err := assertData(data); err != nil {
		return nil, err
	}
	img := image.NewRGBA(image.Rect(0, 0, len(data[0]), len(data)))
	for i, line := range data {
		for j, n := range line {
			c, ok := cmap[n]
			if !ok {
				return nil, fmt.Errorf("color not found: %d", n)
			}
			img.Set(j, i, c)
		}
	}
	return img, nil
}
