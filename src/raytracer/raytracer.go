package main

import (
	"fmt"
	"github.com/SamGostling/RayTracer/internal"
	"image/color"
	"sync"
	"time"
)

func main() {
	createImageWithSphere(
		internal.Sphere{
			Center: internal.Vector{X: -3, Y: 0, Z: -16},
			Radius: 2},
		400,
		500).
		Save(fmt.Sprintf("./renders/sphere%d.png", time.Now().Unix()))
}

func renderGradient() {
	height := 192
	width := 256
	image := internal.NewImage(width, height)
	image.PixelTransform(func(x, y int) color.RGBA {
		v := internal.Vector{X: float64(y) / float64(height), Y: float64(x) / float64(width), Z: 0.2}

		return color.RGBA{
			R: uint8(255 * v.X),
			G: uint8(255 * v.Y),
			B: uint8(255 * v.Z),
			A: 255,
		}
	})
	image.Save(fmt.Sprintf("./renders/gradient%d.png", time.Now().Unix()))
}

func createImageWithSphere(sphere internal.Sphere, height, width int) *internal.Image {
	image := internal.NewImage(width, height)
	var wg sync.WaitGroup

	processRow := func(row int) {
		defer wg.Done()
		y := -((2.0*float64(row)+1)/float64(height) - 1)
		for col := 0; col < width; col++ {
			x := ((2.0*float64(col)+1)/float64(width) - 1) * float64(width) / float64(height)
			dir := internal.Vector{X: x, Y: y, Z: -1}.Normalize()
			ray := internal.Ray{Direction: dir}
			c := ray.Cast(sphere)
			colour := color.RGBA{
				R: uint8(255 * c.X),
				G: uint8(255 * c.Y),
				B: uint8(255 * c.Z),
				A: 255,
			}
			image.SetPixel(row, col, colour)
		}
	}

	for row := 0; row < height; row++ {
		wg.Add(1)
		go processRow(row)
	}

	wg.Wait()
	return image
}
