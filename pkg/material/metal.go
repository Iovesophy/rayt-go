package material

import (
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/golang/geo/r3"
)

var (
	_ geometry.Material = (*Metal)(nil)
)

type Metal struct {
	Albedo r3.Vector
}

func NewMetal(albedo r3.Vector) Metal {
	return Metal{Albedo: albedo}
}

func (metal Metal) Scatter(rayIn ray.VertexPair, record *geometry.Record, attenuation *r3.Vector, scattered *ray.VertexPair) bool {
	reflected := reflect(rayIn.Direction.Normalize(), record.Normal)
	*scattered = ray.New(record.VertexP, reflected)
	*attenuation = metal.Albedo
	return scattered.Direction.Dot(record.Normal) > 0
}

func reflect(v, n r3.Vector) r3.Vector {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}
