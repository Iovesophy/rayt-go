package material

import (
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/golang/geo/r3"
)

var (
	_ geometry.Material = (*Lambertian)(nil)
)

type Lambertian struct {
	Albedo r3.Vector
}

func NewLambertian(albedo r3.Vector) Lambertian {
	return Lambertian{Albedo: albedo}
}

func (lambertian Lambertian) Scatter(rayIn ray.VertexPair, record *geometry.Record, attenuation *r3.Vector, scattered *ray.VertexPair) bool {
	// diffuse reflection
	target := record.VertexP.Add(record.Normal).Add(geometry.RandomInUnitSphere())
	*scattered = ray.New(record.VertexP, target.Sub(record.VertexP))
	*attenuation = lambertian.Albedo
	return true
}
