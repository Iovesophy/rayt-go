package main

import (
	"strconv"

	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/image"
	"github.com/Iovesophy/rayt-go/pkg/material"
	"github.com/Iovesophy/rayt-go/pkg/scene"
	"github.com/Iovesophy/rayt-go/pkg/scene/camera"

	"github.com/golang/geo/r3"
)

func Frame(frame int, i int) {
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
		geometry.NewSphere(r3.Vector{X: 0, Y: -0.1 * float64(i), Z: -1.5}, 0.5, material.NewLambertian(scene.NewVector(0.8, 0.3, 0.3))),
		geometry.NewSphere(r3.Vector{X: 0, Y: 100.5, Z: -1.5}, 100, material.NewLambertian(scene.NewVector(0.8, 0.8, 0))),
	)
	result := e.CreateP3Data()
	e.CreateFile("pingpong"+strconv.Itoa(frame)+".ppm", result.Header, result.Body.String())
}

func main() {
	count := 3
	start := 1
	// pingpong
	for i := start; i < count+start; i++ {
		Frame(i, i)
		Frame(i+count, (count+1)-i)
	}
}
