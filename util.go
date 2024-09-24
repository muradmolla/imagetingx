package imagetingx

import (
	"fmt"
	"math"
)

type ColorMap struct {
	R [256]int
	G [256]int
	B [256]int
	A [256]int
}

func (i *Image) ApplyColormMap(cmp ColorMap) {
	if i.img == nil {
		return
	}
	fmt.Printf("Image Bounds: %dx%d", (*i.img).Bounds().Dx(), (*i.img).Bounds().Dy())
}

func clamp(i, min, max float64) float64 {
	return math.Min(math.Max(i, min), max)
}

func scale(c uint8, scaler float64) uint8 {
	return uint8(clamp(float64(c)*scaler, 0, 255))
}
