package main

import (
	"fmt"
	"github.com/SamGostling/RayTracer/internal"
	"image/color"
	"time"
)

func main() {
	height := 192
	width := 256
	image := internal.NewImage(width, height)
	image.PixelTransform(func(x, y int) color.RGBA {
		v := internal.Vector{X: float64(y) / float64(height), Y: float64(x) / float64(width), Z: 0.2}
		v = v.AddScalar(1)
		return color.RGBA{
			R: uint8(255 * v.X),
			G: uint8(255 * v.Y),
			B: uint8(255 * v.Z),
			A: 255,
		}
	})
	image.Save(fmt.Sprintf("test%d.png", time.Now().Unix()))
}
