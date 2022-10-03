package material

import (
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/Iovesophy/rayt-go/pkg/scene"
	"github.com/golang/geo/r3"
)

var (
	_ geometry.Material = (*Normal)(nil)
)

type Normal struct {
	Bright float64
}

func NewNormal(bright float64) Normal {
	return Normal{Bright: bright}
}

func (normal Normal) Scatter(rayIn ray.VertexPair, record *geometry.Record, attenuation *r3.Vector, scattered *ray.VertexPair) bool {
	*attenuation = scene.NewVector(record.Normal.X, record.Normal.Y, record.Normal.Z).Add(scene.NewVector(1.0, 1.0, 1.0)).Mul(normal.Bright)
	*scattered = ray.VertexPair{}
	return true
}
