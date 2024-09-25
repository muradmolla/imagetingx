package imagetingx_test

import (
	"image"
	"image/color"
	"math/rand"
	"testing"

	"github.com/muradmolla/imagetingx"
)

func TestGamma(t *testing.T) {
	// Testing both gamma functions match
	img := createImage(increasingColor)
	img.MapGamma(2.0)
	imgT := createImage(increasingColor)
	imgT.BruteGamma(2.0)
	r := compareImages(img, imgT)
	if !r {
		t.Errorf("Images does not match when gamma")
	}
	imgT = createImage(increasingColor)
	imgT.Gamma(2.0)
	r = compareImages(img, imgT)
	if !r {
		t.Errorf("Images does not match when gamma")
	}
}

func increasingColor(x, y int) color.NRGBA {
	return color.NRGBA{R: uint8(x % 255), G: uint8(y % 255), B: uint8((x + y) % 255), A: 255}
}

func randomColor(int, int) color.NRGBA {
	rndFn := func() uint8 {
		return uint8(rand.Float32() * 255)
	}

	return color.NRGBA{R: rndFn(), G: rndFn(), B: rndFn(), A: 255}
}

func createImage(imgFunc func(int, int) color.NRGBA) imagetingx.ImgX {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, imgFunc(x, y))
		}
	}

	result, err := imagetingx.New(img)
	if err != nil {
		panic(err)
	}

	return *result
}
