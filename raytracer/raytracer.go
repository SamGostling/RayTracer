package main

import (
	"fmt"
	"github.com/SamGostling/RayTracer/camera"
	"github.com/SamGostling/RayTracer/material"
	"github.com/SamGostling/RayTracer/render"
	"github.com/SamGostling/RayTracer/shape"
	"github.com/SamGostling/RayTracer/vector"
	"image/color"
	"sync"
	"time"
)

func main() {
	var spheres = []shape.Sphere{
		{
			Center:   vector.Vector{X: -3, Y: 0, Z: -16},
			Radius:   2,
			Material: material.ShinyYellow,
		},
		{
			Center:   vector.Vector{X: -1.0, Y: -1.5, Z: -12},
			Radius:   2,
			Material: material.BlueMetal,
		},
		{
			Center:   vector.Vector{X: 1.5, Y: -0.5, Z: -18},
			Radius:   3,
			Material: material.GreenRubber,
		},
		{
			Center:   vector.Vector{X: 7, Y: 5, Z: -18},
			Radius:   4,
			Material: material.RedPlastic,
		},
		{
			Center:   vector.Vector{X: 5, Y: -5, Z: -18},
			Radius:   2,
			Material: material.WhiteMarble,
		},
	}

	var lights = []shape.Light{
		shape.NewLight(vector.Vector{X: -20, Y: 20, Z: 20}, 1.0),
		shape.NewLight(vector.Vector{X: 30, Y: 50, Z: -25}, 0.9),
		shape.NewLight(vector.Vector{X: 30, Y: 20, Z: 30}, 0.8),
	}
	cam := camera.NewCamera(800)
	createImageWithSphere(
		spheres,
		lights,
		cam).
		Save(fmt.Sprintf("./renders/sphere%d.png", time.Now().Unix()))
}

func renderGradient() {
	height := 192
	width := 256
	image := render.NewImage(width, height)
	image.PixelTransform(func(x, y int) color.RGBA {
		v := vector.Vector{X: float64(y) / float64(height), Y: float64(x) / float64(width), Z: 0.2}

		return color.RGBA{
			R: uint8(255 * v.X),
			G: uint8(255 * v.Y),
			B: uint8(255 * v.Z),
			A: 255,
		}
	})
	image.Save(fmt.Sprintf("./renders/gradient%d.png", time.Now().Unix()))
}

func createImageWithSphere(spheres []shape.Sphere, lights []shape.Light, cam camera.Camera) *render.Image {
	image := render.NewImage(cam.Width(), cam.Height())
	var wg sync.WaitGroup

	processRow := func(row int) {
		defer wg.Done()
		y := -((2.0*float64(row)+1)/float64(cam.Height()) - 1)
		for col := 0; col < cam.Width(); col++ {
			x := ((2.0*float64(col)+1)/float64(cam.Width()) - 1) * float64(cam.Width()) / float64(cam.Height())
			dir := vector.Vector{X: x, Y: y, Z: -1}.Normalize()
			ray := vector.Ray{Direction: dir}
			scene := render.Scene{Spheres: spheres, Lights: lights}
			c := scene.CastRay(ray).Color
			colour := color.RGBA{
				R: uint8(255 * c.X),
				G: uint8(255 * c.Y),
				B: uint8(255 * c.Z),
				A: 255,
			}
			image.SetPixel(col, row, colour)
		}
	}

	for row := 0; row < cam.Height(); row++ {
		wg.Add(1)
		go processRow(row)
	}

	wg.Wait()
	return image
}
