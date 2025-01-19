package internal

type Ray struct {
	Origin, Direction Vector
}

/*
Cast casts a ray to a sphere and returns the color of the intersection.
*/
func (r Ray) Cast(sphere Sphere) Vector {
	hits, _ := sphere.QuadIntersect(r)
	if hits {
		return Vector{0.2, 0.7, 0.8} // red
	}

	return Vector{0.4, 0.4, 0.3}
}
