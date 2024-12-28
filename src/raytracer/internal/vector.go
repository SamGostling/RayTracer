package internal

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X*+v.Y*v.Y + v.Z*v.Z)
}
func (v Vector) Dot(i Vector) float64 {
	return v.X*i.X + v.Y*i.Y + v.Z*i.Z
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector) MultiplyScalar(t float64) Vector {
	return Vector{v.X * t, v.Y * t, v.Z * t}
}

func (v Vector) AddScalar(t float64) Vector {
	return Vector{v.X + t, v.Y + t, v.Z + t}
}
