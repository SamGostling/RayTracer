package internal

type Ray struct {
	Origin, Direction Vector
}

/*
Cast casts a ray to a sphere and returns the color of the intersection.
*/
func (r Ray) Cast(scene Scene) Material {
	return scene.SceneIntersect(r)
}
