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
		return sphere.Material.Color // red
	}

	return Background
}
