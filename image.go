package imagetingx

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type Image struct {
	input *os.File
	img   *image.NRGBA
}

func New(input string) (*Image, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error while opening file. error: %s", err)
	}
	src, err := jpeg.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error while decoding jpeg. error: %s", err)
	}

	img := image.NewNRGBA(src.Bounds())
	draw.Draw(img, img.Bounds(), src, img.Bounds().Min, draw.Src)

	return &Image{input: file, img: img}, nil
}

func (img *Image) Save(output string) error {
	defer img.input.Close()
	fs, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error while creating output. error: %s", err)
	}

	var result image.Image = img.img
	jpeg.Encode(fs, result, nil)

	return nil
}
