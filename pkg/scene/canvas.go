package scene

import (
	"rayt-go/pkg/ray"
	"rayt-go/pkg/sphere"

	"github.com/golang/geo/r3"
)

var center = r3.Vector{
	X: 0.0,
	Y: 0.0,
	Z: -1.0,
}
var unitVector = r3.Vector{
	X: 1.0,
	Y: 1.0,
	Z: 1.0,
}

func Pixel(r ray.VertexPair) r3.Vector {
	sphereSize := 0.6
	t := sphere.Hit(center, sphereSize, r)
	if t > 0.0 {
		n := PointAtParameter(t, r).Sub(center).Normalize()
		result := r3.Vector{
			X: n.X,
			Y: n.Y,
			Z: n.Z,
		}.Add(
			unitVector,
		).Mul(
			0.5,
		)
		return result
	}
	return Background(r)
}

func Background(r ray.VertexPair) r3.Vector {
	unitVector := r3.Vector{
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}
	unit := r.Direction.Normalize()
	t := 0.5*unit.Y + 1.0
	result := unitVector.Add(
		r3.Vector{
			X: 0.7,
			Y: 0.7,
			Z: 0.7,
		}.Mul(
			t,
		))
	return result
}
