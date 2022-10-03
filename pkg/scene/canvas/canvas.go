package canvas

import (
	"math"

	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/prand"
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

func RandomInUnitSphere() r3.Vector {
	var p = scene.NewVector(0, 0, 0)
	for {
		randfloat64X, err := prand.Float64()
		if err != nil {
			panic(err)
		}
		randfloat64Y, err := prand.Float64()
		if err != nil {
			panic(err)
		}
		randfloat64Z, err := prand.Float64()
		if err != nil {
			panic(err)
		}
		p = scene.NewVector(randfloat64X, randfloat64Y, randfloat64Z).Mul(2).Sub(unitVector)
		if p.Norm2() < 1 {
			break
		}
	}
	return p
}

func (color Color) Pixel(vertexpair ray.VertexPair, world geometry.Hitable) r3.Vector {
	var record geometry.Record
	if world.Hit(vertexpair, 0.001, math.MaxFloat64, &record) {
		// diffuse reflection
		target := record.VertexP.Add(record.Normal).Add(RandomInUnitSphere())
		return color.Pixel(ray.New(record.VertexP, target.Sub(record.VertexP)), world).Mul(0.5)
		// color mapping
		//return scene.NewVector(record.Normal.X, record.Normal.Y, record.Normal.Z).Add(unitVector).Mul(0.5)
	}
	return color.Background(vertexpair)
}
