package sphere

import (
	"rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

func Hit(center r3.Vector, radius float64, ray scene.Ray) float64 {
	oc := ray.Origin.Sub(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - radius*radius
	if Detector(a, b, c) < 0 {
		return -1.0
	}
	return Solver(a, b, c)
}
