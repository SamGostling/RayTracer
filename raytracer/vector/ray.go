package vector

type Ray struct {
	Origin, Direction Vector
}

/*PointAtParameter returns the point at a distance along the ray.*/
func (r Ray) PointAtParameter(dist float64) Vector {
	return r.Origin.Add(r.Direction.MultiplyScalar(dist))
}
