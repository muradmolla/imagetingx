package imagetingx

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type ImgX struct {
	img *image.NRGBA
}

func New(input interface{}) (*ImgX, error) {
	var img *image.NRGBA
	var err error
	switch v := input.(type) {
	case string:
		img, err = newFile(v)
		if err != nil {
			return nil, err
		}
	case *image.NRGBA:
		img = v
	case *image.Image:
		img = image.NewNRGBA((*v).Bounds())
		draw.Draw(img, img.Bounds(), *v, img.Bounds().Min, draw.Src)
	case image.NRGBA:
		img = &v
	case image.Image:
		img = image.NewNRGBA(v.Bounds())
		draw.Draw(img, img.Bounds(), v, img.Bounds().Min, draw.Src)
	default:
		return nil, fmt.Errorf("invalid input type. Valid types are string, *image.NRGBA, *image.Image, image.NRGBA, image.Image")
	}
	return &ImgX{img: img}, nil
}

func newFile(input string) (*image.NRGBA, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error while opening file. error: %s", err)
	}

	src, err := jpeg.Decode(file)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error while decoding jpeg. error: %s", err)
	}

	img := image.NewNRGBA(src.Bounds())
	draw.Draw(img, img.Bounds(), src, img.Bounds().Min, draw.Src)

	return img, nil
}

func (img *ImgX) Save(output string) error {
	fs, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error while creating output. error: %s", err)
	}

	var result image.Image = img.img
	jpeg.Encode(fs, result, nil)

	return nil
}
