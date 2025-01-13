package internal

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	Width, Height int
	Img           *image.RGBA
}

func NewImage(width int, height int) *Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return &Image{width, height, img}
}

func (i *Image) PixelTransform(transform func(int, int) color.RGBA) {
	for x := 0; x < i.Width; x++ {
		for y := 0; y < i.Height; y++ {
			i.SetPixel(x, y, transform(x, y))
		}
	}
}

func (i *Image) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, i.Img)
}

func (i *Image) SetPixel(x int, y int, color color.RGBA) {
	i.Img.Set(x, y, color)
}
