package camera

import (
	"math"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/utils"

	"github.com/golang/geo/r3"
)

type Parts struct {
	Origin          r3.Vector
	LowerLeftCorner r3.Vector
	Horizontal      r3.Vector
	Vertical        r3.Vector
}

func New(origin, lowerLeftCorner, horizontal, vertical r3.Vector) Parts {
	return Parts{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
	}
}

func Main(X, Y int) (r3.Vector, r3.Vector, r3.Vector, r3.Vector) {
	gcd := utils.Gcd(uint64(X), uint64(Y))
	x := float64(X) / float64(gcd)
	y := float64(Y) / float64(gcd)
	origin := scene.NewVector(
		0.0,
		0.0,
		0.0,
	)
	lowerLeftCorner := r3.Vector{
		X: -x,
		Y: -y,
		Z: -1.0,
	}
	// lowerLeftCorner から基底ベクトルを求める
	horizontal := scene.NewVector(
		math.Abs(lowerLeftCorner.X*2.0),
		0.0,
		0.0,
	)
	vertical := scene.NewVector(
		0.0,
		math.Abs(lowerLeftCorner.Y*2.0),
		0.0,
	)
	return origin, lowerLeftCorner, horizontal, vertical
}
