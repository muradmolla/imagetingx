package imagetingx_test

import (
	"github.com/muradmolla/imagetingx"
)

func compareImages(img1, img2 imagetingx.ImgX) bool {
	i1 := img1.Image()
	i2 := img2.Image()
	if i1.Bounds() != i2.Bounds() {
		return false
	}
	for x := i1.Bounds().Min.X; x < i1.Bounds().Max.X; x++ {
		for y := i1.Bounds().Min.Y; y < i1.Bounds().Max.Y; y++ {
			if i1.At(x, y) != i2.At(x, y) {
				return false
			}
		}
	}
	return true
}
