package camera

import "github.com/Iovesophy/rayt-go/pkg/ray"

func (camera Parts) Ray(h, v float64) ray.VertexPair {
	return ray.New(
		camera.Origin,
		camera.LowerLeftCorner.Add(
			camera.Horizontal.Mul(h),
		).Add(
			camera.Vertical.Mul(v),
		).Sub(camera.Origin),
	)
}
