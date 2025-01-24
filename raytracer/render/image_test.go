package render

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	width, height := 256, 192
	img := NewImage(width, height)
	assert.Equal(t, width, img.Width, "Width should be equal")
	assert.Equal(t, height, img.Height, "Height should be equal")
	assert.NotNil(t, img.Img, "Image should not be nil")
}

func TestPixelTransform(t *testing.T) {
	width, height := 2, 2
	img := NewImage(width, height)
	transform := func(x, y int) color.RGBA {
		return color.RGBA{uint8(x * 255), uint8(y * 255), 0, 255}
	}
	img.PixelTransform(transform)
	assert.Equal(t, color.RGBA{0, 0, 0, 255}, img.Img.At(0, 0), "Pixel (0,0) should be black")
	assert.Equal(t, color.RGBA{255, 0, 0, 255}, img.Img.At(1, 0), "Pixel (1,0) should be red")
	assert.Equal(t, color.RGBA{0, 255, 0, 255}, img.Img.At(0, 1), "Pixel (0,1) should be green")
	assert.Equal(t, color.RGBA{255, 255, 0, 255}, img.Img.At(1, 1), "Pixel (1,1) should be yellow")
}

func TestSave(t *testing.T) {
	width, height := 2, 2
	img := NewImage(width, height)
	transform := func(x, y int) color.RGBA {
		return color.RGBA{uint8(x * 255), uint8(y * 255), 0, 255}
	}
	img.PixelTransform(transform)
	filename := "./test_image.png"
	img.Save(filename)
	defer os.Remove(filename)

	_, err := os.Stat(filename)
	assert.False(t, os.IsNotExist(err), "File should exist")
}
