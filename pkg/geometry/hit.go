package geometry

import (
	"github.com/Iovesophy/rayt-go/pkg/prand"
	"github.com/Iovesophy/rayt-go/pkg/ray"
	"github.com/Iovesophy/rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

var (
	_ Hitable = (*World)(nil)
	_ Hitable = (*Sphere)(nil)
)

type Hitable interface {
	Hit(vertexpair ray.VertexPair, min float64, max float64, record *Record) bool
}

type Material interface {
	Scatter(rayIn ray.VertexPair, record *Record, attenuation *r3.Vector, scattered *ray.VertexPair) bool
}

type Record struct {
	VertexT  float64
	VertexP  r3.Vector
	Normal   r3.Vector
	Material *Material
}

type World struct {
	List []Hitable
}

func New(hitable []Hitable) World {
	return World{List: hitable}
}

type Sphere struct {
	Center   r3.Vector
	Radius   float64
	Material Material
}

func NewSphere(center r3.Vector, radius float64, material Material) Sphere {
	return Sphere{Center: center, Radius: radius, Material: material}
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
		p = scene.NewVector(randfloat64X, randfloat64Y, randfloat64Z).Mul(2).Sub(scene.UnitVector)
		if p.Norm2() < 1 {
			break
		}
	}
	return p
}

func (world World) Hit(vertexpair ray.VertexPair, min float64, max float64, record *Record) bool {
	var memory Record
	hitted := false
	far := max
	for i := 0; i < len(world.List); i++ {
		if world.List[i].Hit(vertexpair, min, far, &memory) {
			hitted = true
			far = memory.VertexT
			*record = memory
		}
	}
	return hitted
}

func (sphere Sphere) Hit(vertexpair ray.VertexPair, min float64, max float64, record *Record) bool {
	oc := vertexpair.Origin.Sub(sphere.Center)
	a := vertexpair.Direction.Dot(vertexpair.Direction)
	b := 2.0 * oc.Dot(vertexpair.Direction)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius
	if SphereDetector(a, b, c) > 0 {
		Vertex1 := SphereNegSolver(a, b, c)
		if Vertex1 > min && Vertex1 < max {
			record.VertexT = Vertex1
			record.VertexP = ray.PointAtParameter(record.VertexT, vertexpair)
			record.Normal = record.VertexP.Sub(sphere.Center).Mul(1.0 / sphere.Radius)
			record.Material = &sphere.Material
			return true
		}
		Vertex2 := SpherePosSolver(a, b, c)
		if Vertex2 > min && Vertex2 < max {
			record.VertexT = Vertex2
			record.VertexP = ray.PointAtParameter(record.VertexT, vertexpair)
			record.Normal = record.VertexP.Sub(sphere.Center).Mul(1.0 / sphere.Radius)
			record.Material = &sphere.Material
			return true
		}
	}
	return false
}
