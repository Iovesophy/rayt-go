package canvas

import (
	"math"

	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/Iovesophy/rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

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
			return scene.MulVector(attenuation, color.Pixel(scattered, world, depth+1))
		}
		return scene.NewVector(0, 0, 0)
	}
	return color.Background(vertexpair)
}
