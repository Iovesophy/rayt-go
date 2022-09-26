package scene

import (
	"github.com/golang/geo/r3"
)

// 線形補完パラメータ t で補完している
func PointAtParameter(t float64, ray Ray) r3.Vector {
	return r3.Vector{
		X: ray.Origin.X + t*ray.Direction.X,
		Y: ray.Origin.Y + t*ray.Direction.Y,
		Z: ray.Origin.Z + t*ray.Direction.Z,
	}
}
