package main

import "fmt"

type Image struct {
	F string
	X int
	Y int
	L int
	R int
	G int
	B int
}

func (image Image) Header() {
	fmt.Printf("%s\n%d %d\n%d\n", image.F, image.X, image.Y, image.L)
}

func (image Image) Data() Image {
	for j := 0; j < image.Y; j++ {
		for i := 0; i < image.X; i++ {
			r := float64(i) / float64(image.X)
			g := float64(j) / float64(image.Y)
			b := 0.2
			image.R = int(255.99 * r)
			image.G = int(255.99 * g)
			image.B = int(255.99 * b)
			fmt.Printf("%d %d %d\n", image.R, image.G, image.B)
		}
	}
	return image
}

func main() {
	image := Image{}
	image.F = "P3"
	image.X = 200
	image.Y = 100
	image.L = 255
	image.Header()
	image.Data()
}
