package ray

import (
	"rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

type VertexPair struct {
	Origin    r3.Vector
	Direction r3.Vector
}

func New(origin r3.Vector, direction r3.Vector) VertexPair {
	return VertexPair{Origin: origin, Direction: direction}
}

func PointAtParameter(t float64, vertexpair VertexPair) r3.Vector {
	return scene.NewVector(
		vertexpair.Origin.X+t*vertexpair.Direction.X,
		vertexpair.Origin.Y+t*vertexpair.Direction.Y,
		vertexpair.Origin.Z+t*vertexpair.Direction.Z,
	)
}
