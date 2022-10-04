package scene

import "github.com/golang/geo/r3"

var UnitVector = NewVector(1.0, 1.0, 1.0)

func NewVector(x, y, z float64) r3.Vector {
	return r3.Vector{X: x, Y: y, Z: z}
}

func MulVector(a, b r3.Vector) r3.Vector {
	return NewVector(a.X*b.X, a.Y*b.Y, a.Z*b.Z)
}
