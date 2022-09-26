package sphere

import (
	"rayt-go/pkg/ray"

	"github.com/golang/geo/r3"
)

func Hit(center r3.Vector, radius float64, r ray.VertexPair) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	if Detector(a, b, c) < 0 {
		return -1.0
	}
	return Solver(a, b, c)
}
