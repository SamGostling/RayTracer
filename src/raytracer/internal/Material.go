package internal

type Material struct {
	Color             Vector
	Albedo            Vector
	SpecularComponent float32
}

func NewMaterial(color Vector, albedo Vector, specularComponent float32) Material {
	return Material{
		Color:             color,
		Albedo:            albedo,
		SpecularComponent: specularComponent,
	}
}

var (
	Background  = NewMaterial(Vector{0.02, 0.02, 0.02}, Vector{0.0, 0.0, 0.0}, 0.0)
	ShinyYellow = NewMaterial(Vector{0.95, 0.95, 0.4}, Vector{0.7, 0.6, 0}, 30.0)
	GreenRubber = NewMaterial(Vector{0.3, 0.7, 0.3}, Vector{0.9, 0.1, 0}, 1.0)
	BlueMetal   = NewMaterial(Vector{0.1, 0.1, 0.8}, Vector{0.6, 0.6, 0.8}, 50.0)
	RedPlastic  = NewMaterial(Vector{0.8, 0.1, 0.1}, Vector{0.7, 0.3, 0.1}, 10.0)
	WhiteMarble = NewMaterial(Vector{0.9, 0.9, 0.9}, Vector{0.8, 0.8, 0.8}, 5.0)
)
