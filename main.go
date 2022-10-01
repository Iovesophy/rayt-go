package main

import (
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/image"
	"rayt-go/pkg/scene/camera"

	"github.com/golang/geo/r3"
)

func main() {
	//defer profile.Start(profile.ProfilePath("./assets")).Stop()
	e := image.Elements{}
	e.Format = "P3"
	e.X = 800
	e.Y = 400
	e.Sampling = 100
	e.MaxBright = 255
	e.Camera = camera.New(camera.Main(e.X, e.Y))
	e.World = append(e.World,
		geometry.NewSphere(r3.Vector{X: 0, Y: 0, Z: -1.5}, 0.925),
		geometry.NewSphere(r3.Vector{X: 0, Y: 100.925, Z: -1.5}, 100),
	)
	result := e.CreateP3Data()
	e.CreateFile("test.ppm", result.Header, result.Body.String())
}
