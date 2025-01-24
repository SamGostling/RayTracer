package material

import "github.com/SamGostling/RayTracer/vector"

type Material struct {
	Color             vector.Vector
	Albedo            vector.Vector
	SpecularComponent float32
}

func NewMaterial(color vector.Vector, albedo vector.Vector, specularComponent float32) Material {
	return Material{
		Color:             color,
		Albedo:            albedo,
		SpecularComponent: specularComponent,
	}
}

var (
	Background  = NewMaterial(vector.Vector{X: 0.02, Y: 0.02, Z: 0.02}, vector.Vector{}, 0.0)
	ShinyYellow = NewMaterial(vector.Vector{X: 0.95, Y: 0.95, Z: 0.4}, vector.Vector{X: 0.7, Y: 0.6}, 30.0)
	GreenRubber = NewMaterial(vector.Vector{X: 0.3, Y: 0.7, Z: 0.3}, vector.Vector{X: 0.9, Y: 0.1}, 1.0)
	BlueMetal   = NewMaterial(vector.Vector{X: 0.1, Y: 0.1, Z: 0.8}, vector.Vector{X: 0.6, Y: 0.6, Z: 0.8}, 50.0)
	RedPlastic  = NewMaterial(vector.Vector{X: 0.8, Y: 0.1, Z: 0.1}, vector.Vector{X: 0.7, Y: 0.3, Z: 0.1}, 10.0)
	WhiteMarble = NewMaterial(vector.Vector{X: 0.9, Y: 0.9, Z: 0.9}, vector.Vector{X: 0.8, Y: 0.8, Z: 0.8}, 5.0)
)
