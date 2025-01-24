package render

import (
	"github.com/SamGostling/RayTracer/material"
	"github.com/SamGostling/RayTracer/shape"
	"github.com/SamGostling/RayTracer/vector"
	"math"
)

type Scene struct {
	Spheres            []shape.Sphere
	EfficientIntersect bool
}

func (s *Scene) SceneIntersect(ray vector.Ray) material.Material {
	spheresDist := math.MaxFloat64
	intersectMaterial := material.Background

	for _, obj := range s.Spheres {
		intersect, dist := s.intersectSphere(ray, obj)

		if intersect && dist < spheresDist {
			spheresDist = dist
			intersectMaterial = obj.Material
		}
	}
	return intersectMaterial
}

func (s *Scene) intersectSphere(ray vector.Ray, sphere shape.Sphere) (bool, float64) {
	if s.EfficientIntersect {
		return sphere.GeoIntersect(ray)
	}
	return sphere.QuadIntersect(ray)
}
