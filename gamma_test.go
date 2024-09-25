package imagetingx_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/muradmolla/imagetingx"
)

func TestGamma(t *testing.T) {
	img, err := imagetingx.New("test.jpg")
	if err != nil {
		t.Fatalf("cannot init image. error :%s", err)
	}
	img.Gamma(0.5)
	img.Save("test_gamma.jpg")
}

func createImage() imagetingx.ImgX {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	return imagetingx.ImgX{img: img}
}
