package image

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/Iovesophy/rayt-go/pkg/format"
	"github.com/Iovesophy/rayt-go/pkg/geometry"
	"github.com/Iovesophy/rayt-go/pkg/prand"
	"github.com/Iovesophy/rayt-go/pkg/scene"
	"github.com/Iovesophy/rayt-go/pkg/scene/camera"
	"github.com/Iovesophy/rayt-go/pkg/scene/canvas"

	"github.com/golang/geo/r3"
)

type RGB struct {
	R int
	G int
	B int
}

type Elements struct {
	Format    string
	X         int
	Y         int
	Sampling  int
	MaxBright int
	Header    string
	Body      strings.Builder
	Color     RGB
	Camera    camera.Parts
	World     []geometry.Hitable
	Depth     int
	MaxDepth  int
}

func (img Elements) CreateHeader() string {
	return fmt.Sprintf(format.Header, img.Format, img.X, img.Y, img.MaxBright)
}

func (img Elements) CreateP3Data() Elements {
	img.Header = img.CreateHeader()
	wg := new(sync.WaitGroup)
	p := make([]string, img.X*img.Y)
	wg.Add(img.X * img.Y)
	for j := 0; j < img.Y; j++ {
		for i := 0; i < img.X; i++ {
			go Render(i, j, img.Camera, wg, &img, p)
		}
	}
	wg.Wait()
	// transpose
	for j := img.Y - 1; j >= 0; j-- {
		for i := 0; i < img.X; i++ {
			img.Body.WriteString(p[i+j*img.X])
		}
	}
	return img
}

func Render(i int, j int, camera camera.Parts, wg *sync.WaitGroup, img *Elements, p []string) {
	defer wg.Done()
	color := scene.NewVector(0, 0, 0)
	bgcolor := canvas.Color{X: 0.5, Y: 0.7, Z: 1.0}
	SuperSampling(i, j, camera, &color, geometry.New(img.World), bgcolor, img)
	// fix gamma
	color = scene.NewVector(math.Sqrt(color.X), math.Sqrt(color.Y), math.Sqrt(color.Z))
	SetRGB(color, img)
	p[i+j*img.X] = fmt.Sprintf(format.Body, img.Color.R, img.Color.G, img.Color.B)
}

func SuperSampling(i int, j int, camera camera.Parts, color *r3.Vector, world geometry.World, canvas canvas.Color, img *Elements) {
	for s := 0; s < img.Sampling; s++ {
		randfloat64A, err := prand.Float64()
		if err != nil {
			panic(err)
		}
		h := (float64(i) + randfloat64A) / float64(img.X)
		randfloat64B, err := prand.Float64()
		if err != nil {
			panic(err)
		}
		v := (float64(j) + randfloat64B) / float64(img.Y)
		ray := camera.Ray(h, v)
		*color = color.Add(canvas.Pixel(ray, world, img.Depth, img.MaxDepth))
	}
	*color = color.Mul(1.0 / float64(img.Sampling))
}

func SetRGB(color r3.Vector, img *Elements) {
	// R
	img.Color.R = int(255.99 * color.X)

	// G
	img.Color.G = int(255.99 * color.Y)

	// B
	img.Color.B = int(255.99 * color.Z)
}
