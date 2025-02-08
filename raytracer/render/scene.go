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
	var point, N vector.Vector
	intersectMaterial := material.Background

	if !s.sceneIntersect(ray, &intersectMaterial, &point, &N) {
		return intersectMaterial
	}

	diffuseLightIntensity, specularLightIntensity := 0.0, 0.0
	for _, light := range s.Lights {
		lightDir := light.Center().Subtract(point).Normalize()
		diffuseLightIntensity += light.Intensity() * math.Max(0, lightDir.Dot(N))
		specularLightIntensity += math.Pow(
			math.Max(0, -reflect(lightDir.Negate(), N).Dot(ray.Direction)),
			intersectMaterial.SpecularComponent) * light.Intensity()
	}

	materialColor := intersectMaterial.Color.
		MultiplyScalar(diffuseLightIntensity).
		MultiplyScalar(intersectMaterial.Albedo.X).
		Add(vector.Vector{X: 1, Y: 1, Z: 1}.
			MultiplyScalar(specularLightIntensity).MultiplyScalar(intersectMaterial.Albedo.Y))

	// Something seems off in calculating the color of the intersected object
	// clamping color to avoid overflow
	return material.Material{Color: clampColor(materialColor)}
}

/*
sceneIntersect checks if a ray intersects with any of the spheres in the scene and updates the intersected material,
the hit point and the normal at the hit point.
*/
func (s *Scene) sceneIntersect(ray vector.Ray, intersectMaterial *material.Material, hit, normal *vector.Vector) bool {
	spheresDist := math.MaxFloat64

	for _, sphere := range s.Spheres {
		intersect, dist := s.intersectSphere(ray, sphere)

		if intersect && dist < spheresDist {
			spheresDist = dist
			*hit = ray.PointAtParameter(dist)
			*normal = hit.Subtract(sphere.Center).Normalize()
			*intersectMaterial = sphere.Material
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

func reflect(I, N vector.Vector) vector.Vector {
	return I.Subtract(N.MultiplyScalar(2 * I.Dot(N)))
}

func clampColor(color vector.Vector) vector.Vector {
	return vector.Vector{
		X: math.Max(0, math.Min(1, color.X)),
		Y: math.Max(0, math.Min(1, color.Y)),
		Z: math.Max(0, math.Min(1, color.Z)),
	}
}
