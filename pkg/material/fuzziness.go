package material

import (
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/golang/geo/r3"
)

var (
	_ geometry.Material = (*Fuzziness)(nil)
)

type Fuzziness struct {
	Albedo r3.Vector
	Fuzz   float64
}

func NewFuzziness(albedo r3.Vector, fuzz float64) Fuzziness {
	return Fuzziness{Albedo: albedo, Fuzz: fuzz}
}

func (fuzziness Fuzziness) Scatter(rayIn ray.VertexPair, record *geometry.Record, attenuation *r3.Vector, scattered *ray.VertexPair) bool {
	reflected := reflect(rayIn.Direction.Normalize(), record.Normal)
	// recalculate for fuzziness
	reflected = reflected.Add(geometry.RandomInUnitSphere().Mul(fuzziness.Fuzz))
	*scattered = ray.New(record.VertexP, reflected)
	*attenuation = fuzziness.Albedo
	return scattered.Direction.Dot(record.Normal) > 0
}
