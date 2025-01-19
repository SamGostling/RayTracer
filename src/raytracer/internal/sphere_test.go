package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sphere = Sphere{Center: Vector{0, 0, 0}, Radius: 1}

func TestSphereIntersect(t *testing.T) {
	// Ray that intersectsQuad the sphere at two points
	ray1 := Ray{Origin: Vector{0, 0, -3}, Direction: Vector{0, 0, 1}}
	intersectsQuad, distanceQuad := sphere.QuadIntersect(ray1)
	intersectsGeo, distanceGeo := sphere.GeoIntersect(ray1)
	assert.Equal(t, intersectsQuad, intersectsGeo, "QuadIntersect and GeoIntersect should return the same intersection result")
	assert.InDelta(t, distanceQuad, distanceGeo, 1e-6, "QuadIntersect and GeoIntersect should return the same intersection distance")
	assert.True(t, intersectsQuad, "Ray should intersect the sphere")
	assert.InDelta(t, 2.0, distanceQuad, 1e-6, "Intersection distanceQuad should be approximately 2.0")
}

func TestSphereIntersectEdgeCases(t *testing.T) {
	// Ray that just touches the sphere (discriminant = 0)
	ray2 := Ray{Origin: Vector{1, 0, -1}, Direction: Vector{0, 0, 1}}
	intersectsQuad, distanceQuad := sphere.QuadIntersect(ray2)
	intersectsGeo, distanceGeo := sphere.GeoIntersect(ray2)
	assert.Equal(t, intersectsQuad, intersectsGeo, "QuadIntersect and GeoIntersect should return the same intersection result")
	assert.InDelta(t, distanceQuad, distanceGeo, 1e-6, "QuadIntersect and GeoIntersect should return the same intersection distance")
	assert.True(t, intersectsQuad, "Ray should just touch the sphere")
	assert.InDelta(t, 1.0, distanceQuad, 1e-6, "Intersection distanceQuad should be approximately 1.0")
}

func TestSphereIntersectMiss(t *testing.T) {
	// Ray that misses the sphere
	ray3 := Ray{Origin: Vector{0, 0, -3}, Direction: Vector{0, 1, 0}}
	intersectsQuad, distanceQuad := sphere.QuadIntersect(ray3)
	intersectsGeo, distanceGeo := sphere.GeoIntersect(ray3)
	assert.Equal(t, intersectsQuad, intersectsGeo, "QuadIntersect and GeoIntersect should return the same intersection result")
	assert.InDelta(t, distanceQuad, distanceGeo, 1e-6, "QuadIntersect and GeoIntersect should return the same intersection distance")
	assert.False(t, intersectsQuad, "Ray should miss the sphere")
	assert.Equal(t, 0.0, distanceQuad, "Intersection distanceQuad should be 0.0")
}

func TestSphereIntersectInside(t *testing.T) {
	// Ray that starts inside the sphere
	ray4 := Ray{Origin: Vector{0, 0, 0}, Direction: Vector{0, 0, 1}}
	intersectsQuad, distanceQuad := sphere.QuadIntersect(ray4)
	intersectsGeo, distanceGeo := sphere.GeoIntersect(ray4)
	assert.Equal(t, intersectsQuad, intersectsGeo, "QuadIntersect and GeoIntersect should return the same intersection result")
	assert.InDelta(t, distanceQuad, distanceGeo, 1e-6, "QuadIntersect and GeoIntersect should return the same intersection distance")
	assert.True(t, intersectsQuad, "Ray should intersect the sphere from inside")
	assert.InDelta(t, 1.0, distanceQuad, 1e-6, "Intersection distanceQuad should be approximately 1.0")
}
