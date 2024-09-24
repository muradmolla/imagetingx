package imagetingx

import (
	"image/color"
)

func (i *Image) Gamma(gamma float64) {
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
