package imagetingx

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

type Image struct {
	input *os.File
	img   *image.Image
}

func New(input string) (*Image, error) {
	file, err := os.Open(input)
	if err != nil {
		msg := fmt.Sprintf("error while opening file. error: %s", err)
		return nil, errors.New(msg)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		msg := fmt.Sprintf("error while decoding jpeg. error: %s", err)
		return nil, errors.New(msg)
	}
	return &Image{input: file, img: &img}, nil
}

func (img *Image) Save(output string) error {
	defer img.input.Close()
	fs, err := os.Create(output)
	if err != nil {
		msg := fmt.Sprintf("error while creating output. error: %s", err)
		return errors.New(msg)
	}

	jpeg.Encode(fs, *img.img, nil)

	return nil
}
