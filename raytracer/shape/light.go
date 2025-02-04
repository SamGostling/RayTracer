package shape

import "github.com/SamGostling/RayTracer/vector"

type Light struct {
	center    vector.Vector
	intensity float64
}

func NewLight(center vector.Vector, intensity float64) Light {
	return Light{
		center:    center,
		intensity: intensity,
	}
}

func (l Light) Center() vector.Vector {
	return l.center
}

func (l Light) Intensity() float64 {
	return l.intensity
}
