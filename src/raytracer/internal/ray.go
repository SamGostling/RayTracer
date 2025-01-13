package internal

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Cast(sphere Sphere) Vector {
	hits, _ := sphere.Intersect(r)
	if hits {
		return Vector{0.2, 0.7, 0.8} // red
	}

	return Vector{0.4, 0.4, 0.3}
}
