package scene

import (
	"math"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/ray"

	"github.com/golang/geo/r3"
)

type Color struct {
	X float64
	Y float64
	Z float64
}

func (color Color) Pixel(vertexpair ray.VertexPair, world geometry.Hitable) r3.Vector {
	var record geometry.Record
	t := world.Hit(vertexpair, 0, math.MaxFloat64, &record)
	if t {
		return r3.Vector{X: record.Normal.X + 1, Y: record.Normal.Y + 1, Z: record.Normal.Z + 1}.Mul(0.5)
	}
	return color.Background(vertexpair)
}

func (color Color) Background(vertexpair ray.VertexPair) r3.Vector {
	unitVector := r3.Vector{
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}
	unit := vertexpair.Direction.Normalize()
	t := 0.5*unit.Y + 1.0
	result := unitVector.Mul(1.0 - t).Add(
		r3.Vector{
			X: color.X,
			Y: color.Y,
			Z: color.Z,
		}.Mul(
			t,
		))
	return result
}
