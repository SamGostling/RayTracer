package internal

import "math"

type Scene struct {
	Spheres            []Sphere
	EfficientIntersect bool
}

func (s *Scene) SceneIntersect(ray Ray) Material {
	spheresDist := math.MaxFloat64
	material := Background

	for _, sphere := range s.Spheres {
		intersect, dist := s.intersectSphere(ray, sphere)

		if intersect && dist < spheresDist {
			spheresDist = dist
			material = sphere.Material
		}
	}
	return material
}

func (s *Scene) intersectSphere(ray Ray, sphere Sphere) (bool, float64) {
	if s.EfficientIntersect {
		return sphere.GeoIntersect(ray)
	}
	return sphere.QuadIntersect(ray)
}
