package ray

import "github.com/golang/geo/r3"

type VertexPair struct {
	Origin    r3.Vector
	Direction r3.Vector
}

func New(origin r3.Vector, direction r3.Vector) VertexPair {
	return VertexPair{Origin: origin, Direction: direction}
}

func PointAtParameter(t float64, r VertexPair) r3.Vector {
	return r3.Vector{
		X: r.Origin.X + t*r.Direction.X,
		Y: r.Origin.Y + t*r.Direction.Y,
		Z: r.Origin.Z + t*r.Direction.Z,
	}
}
