package main

import (
	"rayt-go/pkg/image"
)

func main() {
	img := image.Elements{}
	img.Format = "P3"
	img.X = 800
	img.Y = 400
	img.MaxBright = 255
	result := img.CreateP3Data()
	filename := "test.ppm"
	img.CreateFile(filename, result.Header, result.Body.String())
}
