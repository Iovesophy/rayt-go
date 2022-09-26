package ray

import "github.com/golang/geo/r3"

type VertexPair struct {
	Origin    r3.Vector
	Direction r3.Vector
}

func New(origin r3.Vector, direction r3.Vector) VertexPair {
	return VertexPair{Origin: origin, Direction: direction}
}
