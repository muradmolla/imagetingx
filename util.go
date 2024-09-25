package imagetingx

import (
	"image/color"
	"math"
)

type ColorMap struct {
	R [256]uint8
	G [256]uint8
	B [256]uint8
	A [256]uint8
}

func (i *ImgX) ApplyColormMap(cmp ColorMap) {
	if i.img == nil {
		return
	}

	bounds := (*i.img).Bounds()
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			c := i.img.NRGBAAt(x, y)

			cn := color.NRGBA{
				R: cmp.R[c.R],
				G: cmp.G[c.G],
				B: cmp.B[c.B],
				A: cmp.A[c.A],
			}
			i.img.Set(x, y, cn)
		}
	}
}

func clamp(i, min, max float64) float64 {
	return math.Min(math.Max(i, min), max)
}

func scale(c uint8, scaler float64) uint8 {
	return uint8(clamp(float64(c)*scaler, 0, 255))
}

func newColorMap() ColorMap {
	passIndex := func(v uint8) uint8 {
		return v
	}
	var colorMap ColorMap
	colorMap.R = scaleMapByIndex([256]uint8{}, passIndex)
	colorMap.G = scaleMapByIndex([256]uint8{}, passIndex)
	colorMap.B = scaleMapByIndex([256]uint8{}, passIndex)
	colorMap.A = scaleMapByIndex([256]uint8{}, passIndex)
	return colorMap
}

func scaleMapByIndex(c [256]uint8, scaleFunction func(uint8) uint8) [256]uint8 {
	for i := range c {
		c[i] = scaleFunction(uint8(i))
	}
	return c
}

func scaleMapByValue(c [256]uint8, scaleFunction func(uint8) uint8) [256]uint8 {
	for i, v := range c {
		c[i] = scaleFunction(v)
	}
	return c
}
