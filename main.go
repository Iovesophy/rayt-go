package main

import (
	"math"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/image"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/scene/camera"
	"rayt-go/pkg/utils"

	"github.com/golang/geo/r3"
)

func MainCamera(img image.Elements) (r3.Vector, r3.Vector, r3.Vector, r3.Vector) {
	gcd := utils.Gcd(uint64(img.X), uint64(img.Y))
	x := float64(img.X) / float64(gcd)
	y := float64(img.Y) / float64(gcd)
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

func main() {
	img := image.Elements{}
	img.Format = "P3"
	img.X = 800
	img.Y = 400
	img.Sampling = 100
	img.MaxBright = 255
	img.Camera = camera.New(MainCamera(img))
	var world []geometry.Hitable
	world = append(world,
		geometry.NewSphere(r3.Vector{X: 0, Y: 0, Z: -1.5}, 0.925),
		geometry.NewSphere(r3.Vector{X: 0, Y: 100.925, Z: -1.5}, 100),
	)
	img.World = world
	result := img.CreateP3Data()
	filename := "test.ppm"
	img.CreateFile(filename, result.Header, result.Body.String())
}
