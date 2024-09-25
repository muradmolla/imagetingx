package imagetingx

import (
	"image/color"
)

func (i *ImgX) Gamma(gamma float64) {
	// apply color map
	colorMap := newColorMap()
	fixedScaler := func(v uint8) uint8 {
		return scale(v, gamma)
	}
	colorMap.R = scaleMapByIndex(colorMap.R, fixedScaler)
	colorMap.G = scaleMapByIndex(colorMap.G, fixedScaler)
	colorMap.B = scaleMapByIndex(colorMap.B, fixedScaler)
	colorMap.A = scaleMapByIndex(colorMap.A, fixedScaler)
	i.ApplyColormMap(colorMap)
}

func (i *ImgX) BruteGamma(gamma float64) {
	// basic

	bounds := (*i.img).Bounds()

	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			c := i.img.NRGBAAt(x, y)

			cn := color.NRGBA{
				R: scale(c.R, gamma),
				G: scale(c.G, gamma),
				B: scale(c.B, gamma),
				A: c.A,
			}
			i.img.Set(x, y, cn)
		}
	}
}
