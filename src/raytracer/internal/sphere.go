package internal

import "math"

type Sphere struct {
	Center Vector
	Radius float64
}

/*
QuadIntersect checks if a ray intersects with a sphere and returns a boolean and the distance of the intersection.
It uses the quadratic method to calculate the intersection and is more accurate than the geometric method.
*/
func (s Sphere) QuadIntersect(ray Ray) (bool, float64) {
	originToCenter := ray.Origin.Subtract(s.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * originToCenter.Dot(ray.Direction)
	c := originToCenter.Dot(originToCenter) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return false, 0
	}

	if discriminant == 0 {
		return true, -b / 2 * a
	}

	sqrtDiscriminant := math.Sqrt(discriminant)
	q := -0.5 * (b + sqrtDiscriminant)
	if b > 0 {
		q = -0.5 * (b - sqrtDiscriminant)
	}

	t0 := q / a
	t1 := c / q

	if t0 > t1 {
		t0 = t1
	}

	if t0 < 0 {
		t0 = t1
		if t0 < 0 {
			return false, 0
		}
	}

	return true, t0
}

/*
GeoIntersect checks if a ray intersects with a sphere and returns a boolean and the distance of the intersection.
It uses the geometric method to calculate the intersection and is more efficient than the quadratic method.
*/
func (s Sphere) GeoIntersect(ray Ray) (bool, float64) {
	originToCenter := s.Center.Subtract(ray.Origin)
	projectionLength := originToCenter.Dot(ray.Direction)
	perpendicularDistanceSquared := originToCenter.Dot(originToCenter) - projectionLength*projectionLength
	if perpendicularDistanceSquared > s.Radius*s.Radius {
		return false, 0
	}
	halfChordLength := math.Sqrt(s.Radius*s.Radius - perpendicularDistanceSquared)
	intersection1 := projectionLength - halfChordLength
	intersection2 := projectionLength + halfChordLength
	if intersection1 < 0 {
		intersection1 = intersection2
	}
	if intersection1 < 0 {
		return false, 0
	}
	return true, intersection1
}
