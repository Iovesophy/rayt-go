package canvas

import (
	"math"

	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/Iovesophy/rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

var unitVector = scene.NewVector(1.0, 1.0, 1.0)

type Color struct {
	X float64
	Y float64
	Z float64
}

func (color Color) Pixel(vertexpair ray.VertexPair, world geometry.Hitable, depth int) r3.Vector {
	var record geometry.Record
	if world.Hit(vertexpair, 0.001, math.MaxFloat64, &record) {
		var attenuation r3.Vector
		var scattered ray.VertexPair
		if depth < 50 && (*record.Material).Scatter(vertexpair, &record, &attenuation, &scattered) {
			return VectorMul(attenuation, color.Pixel(scattered, world, depth+1))
		}
		return scene.NewVector(0, 0, 0)
	}
	return color.Background(vertexpair)
}

func VectorMul(a, b r3.Vector) r3.Vector {
	return scene.NewVector(a.X*b.X, a.Y*b.Y, a.Z*b.Z)
}
