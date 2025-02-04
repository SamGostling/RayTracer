package render

import (
	"github.com/SamGostling/RayTracer/material"
	"github.com/SamGostling/RayTracer/shape"
	"github.com/SamGostling/RayTracer/vector"
	"math"
)

type Scene struct {
	Spheres            []shape.Sphere
	Lights             []shape.Light
	EfficientIntersect bool
}

/*
CastRay casts a ray into the scene and returns the color of the intersected object.
*/
func (s *Scene) CastRay(ray vector.Ray) material.Material {
	var point vector.Vector
	var N vector.Vector
	intersectMaterial := material.Background

	if !s.sceneIntersect(ray, &intersectMaterial, &point, &N) {
		return intersectMaterial
	}

	diffuseLightIntensity := 0.0
	for _, light := range s.Lights {
		lightDir := light.Center().Subtract(point).Normalize()
		diffuseLightIntensity += light.Intensity() * math.Max(lightDir.Dot(N), 0)
	}
	if diffuseLightIntensity > 1 {
		diffuseLightIntensity = 1
	}

	return material.Material{Color: intersectMaterial.Color.MultiplyScalar(diffuseLightIntensity)}
}

/*
sceneIntersect checks if a ray intersects with any of the spheres in the scene and updates the intersected material,
the hit point and the normal at the hit point.
*/
func (s *Scene) sceneIntersect(ray vector.Ray, intersectMaterial *material.Material, hit, N *vector.Vector) bool {
	spheresDist := math.MaxFloat64

	for _, obj := range s.Spheres {
		intersect, dist := s.intersectSphere(ray, obj)

		if intersect && dist < spheresDist {
			spheresDist = dist
			*hit = ray.Origin.Add(ray.Direction.MultiplyScalar(dist))
			*N = hit.Subtract(obj.Center).Normalize()
			*intersectMaterial = obj.Material
		}
	}
	return spheresDist != math.MaxFloat64
}

func (s *Scene) intersectSphere(ray vector.Ray, sphere shape.Sphere) (bool, float64) {
	if s.EfficientIntersect {
		return sphere.GeoIntersect(ray)
	}
	return sphere.QuadIntersect(ray)
}
