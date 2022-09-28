package canvas

import (
	"math"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/ray"
	"rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

var unitVector = scene.NewVector(1.0, 1.0, 1.0)

type Color struct {
	X float64
	Y float64
	Z float64
}

func (color Color) Pixel(vertexpair ray.VertexPair, world geometry.Hitable) r3.Vector {
	var record geometry.Record
	t := world.Hit(vertexpair, 0, math.MaxFloat64, &record)
	if t {
		return scene.NewVector(record.Normal.X, record.Normal.Y, record.Normal.Z).Add(unitVector).Mul(0.5)
	}
	return color.Background(vertexpair)
}
