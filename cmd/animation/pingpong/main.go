package main

import (
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/image"
	"rayt-go/pkg/scene/camera"
	"strconv"

	"github.com/golang/geo/r3"
)

func Frame(frame int, i int) {
	e := image.Elements{}
	e.Format = "P3"
	e.X = 800
	e.Y = 400
	e.Sampling = 100
	e.MaxBright = 255
	e.Camera = camera.New(camera.Main(e.X, e.Y))
	e.World = append(e.World,
		geometry.NewSphere(r3.Vector{X: 0, Y: -0.1 * float64(i), Z: -1.5}, 0.5),
		geometry.NewSphere(r3.Vector{X: 0, Y: 100.5, Z: -1.5}, 100),
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
