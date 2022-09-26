package scene

import (
	"rayt-go/pkg/ray"

	"github.com/golang/geo/r3"
)

// 線形補完パラメータ t で補完している
func PointAtParameter(t float64, r ray.VertexPair) r3.Vector {
	return r3.Vector{
		X: r.Origin.X + t*r.Direction.X,
		Y: r.Origin.Y + t*r.Direction.Y,
		Z: r.Origin.Z + t*r.Direction.Z,
	}
}
