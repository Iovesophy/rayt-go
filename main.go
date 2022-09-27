package main

import (
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/image"

	"github.com/golang/geo/r3"
)

func main() {
	img := image.Elements{}
	img.Format = "P3"
	img.X = 800
	img.Y = 400
	img.MaxBright = 255
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
