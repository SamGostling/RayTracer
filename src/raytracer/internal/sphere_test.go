package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	sphere := Sphere{Center: Vector{0, 0, 0}, Radius: 1}

	// Ray that intersects the sphere at two points
	ray1 := Ray{Origin: Vector{0, 0, -3}, Direction: Vector{0, 0, 1}}
	intersects, distance := sphere.Intersect(ray1)
	assert.True(t, intersects, "Ray should intersect the sphere")
	assert.InDelta(t, 2.0, distance, 1e-6, "Intersection distance should be approximately 2.0")

	// Ray that just touches the sphere (discriminant = 0)
	ray2 := Ray{Origin: Vector{1, 0, -1}, Direction: Vector{0, 0, 1}}
	intersects, distance = sphere.Intersect(ray2)
	assert.True(t, intersects, "Ray should just touch the sphere")
	assert.InDelta(t, 1.0, distance, 1e-6, "Intersection distance should be approximately 1.0")

	// Ray that misses the sphere
	ray3 := Ray{Origin: Vector{0, 0, -3}, Direction: Vector{0, 1, 0}}
	intersects, distance = sphere.Intersect(ray3)
	assert.False(t, intersects, "Ray should miss the sphere")
	assert.Equal(t, 0.0, distance, "Intersection distance should be 0.0")

	// Ray that starts inside the sphere
	ray4 := Ray{Origin: Vector{0, 0, 0}, Direction: Vector{0, 0, 1}}
	intersects, distance = sphere.Intersect(ray4)
	assert.True(t, intersects, "Ray should intersect the sphere from inside")
	assert.InDelta(t, 1.0, distance, 1e-6, "Intersection distance should be approximately 1.0")
}
