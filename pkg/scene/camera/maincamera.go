package camera

import (
	"rayt-go/pkg/ray"

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

func (camera Parts) GetCameraRay(h, v float64) ray.VertexPair {
	return ray.New(
		camera.Origin,
		camera.LowerLeftCorner.Add(
			camera.Horizontal.Mul(h),
		).Add(
			camera.Vertical.Mul(v),
		).Sub(camera.Origin),
	)
}
