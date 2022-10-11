package main

import (
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/image"
	"github.com/Iovesophy/rayt-go/pkg/material"
	"github.com/Iovesophy/rayt-go/pkg/scene"
	"github.com/Iovesophy/rayt-go/pkg/scene/camera"

	"github.com/golang/geo/r3"
)

func main() {
	e := image.Elements{}
	e.Format = "P3"
	e.X = 800
	e.Y = 400
	e.Sampling = 100
	e.MaxBright = 255
	e.MaxDepth = 50
	e.Depth = 0
	e.Camera = camera.New(camera.Main(e.X, e.Y))
	e.World = append(e.World,
		geometry.NewSphere(r3.Vector{X: -0.6, Y: 0, Z: -1}, 0.5, material.NewFuzziness(scene.NewVector(0.8, 0.1, 0.1), 1.0)),
		geometry.NewSphere(r3.Vector{X: 0.6, Y: 0, Z: -1}, 0.5, material.NewFuzziness(scene.NewVector(0.3, 0.8, 0.5), 1.0)),
		geometry.NewSphere(r3.Vector{X: 0, Y: -100.5, Z: -1}, 100, material.NewMetal(scene.NewVector(0.7, 0.7, 0.7))),
	)
	result := e.CreateP3Data()
	e.CreateFile("test.ppm", result.Header, result.Body.String())
}
