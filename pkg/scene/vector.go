package scene

import "github.com/golang/geo/r3"

func NewVector(x, y, z float64) r3.Vector {
	return r3.Vector{X: x, Y: y, Z: z}
}
