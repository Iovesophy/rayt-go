package canvas

import (
	"rayt-go/pkg/ray"
	"rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

func (color Color) Background(vertexpair ray.VertexPair) r3.Vector {
	unit := vertexpair.Direction.Normalize()
	t := 0.5*unit.Y + 1.0
	result := unitVector.Mul(1.0 - t).Add(
		scene.NewVector(
			color.X,
			color.Y,
			color.Z,
		).Mul(
			t,
		))
	return result
}
