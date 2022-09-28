package image

import (
	"fmt"
	"math"
	"math/rand"
	"rayt-go/pkg/format"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/scene/camera"
	"rayt-go/pkg/scene/canvas"
	"strings"
	"sync"

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
}

func (img Elements) CreateHeader() string {
	return fmt.Sprintf(format.Header, img.Format, img.X, img.Y, img.MaxBright)
}

func (img Elements) CreateP3Data() Elements {
	img.Header = img.CreateHeader()
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	p := make([]string, img.X*img.Y)
	wg.Add(img.X * img.Y)
	for j := 0; j < img.Y; j++ {
		for i := 0; i < img.X; i++ {
			go Render(i, j, img.Camera, wg, mu, &img, p)
		}
	}
	wg.Wait()
	for _, v := range p {
		img.Body.WriteString(v)
	}
	return img
}

func Render(i int, j int, camera camera.Parts, wg *sync.WaitGroup, mu *sync.Mutex, img *Elements, p []string) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	color := scene.NewVector(0, 0, 0)
	SuperSampling(i, j, camera, &color, geometry.New(img.World), canvas.Color{X: 0.5, Y: 0.7, Z: 1.0}, img)
	// fix ganma
	color = scene.NewVector(math.Sqrt(color.X), math.Sqrt(color.Y), math.Sqrt(color.Z))
	RGBValidation(color, img)
	p[i+j*img.X] = fmt.Sprintf(format.Body, img.Color.R, img.Color.G, img.Color.B)
}

func SuperSampling(i int, j int, camera camera.Parts, color *r3.Vector, world geometry.World, sky canvas.Color, img *Elements) {
	for s := 0; s < img.Sampling; s++ {
		h := (float64(i) + rand.Float64()) / float64(img.X)
		v := (float64(j) + rand.Float64()) / float64(img.Y)
		ray := camera.Ray(h, v)
		*color = color.Add(sky.Pixel(ray, world))
	}
	*color = color.Mul(1.0 / float64(img.Sampling))
}

func RGBValidation(color r3.Vector, img *Elements) {
	// R
	if int(255.99*color.X) < 256 {
		img.Color.R = int(255.99 * color.X)
	} else if int(255.99*color.X) < 0 {
		img.Color.R = 0
	} else {
		img.Color.R = 255
	}

	// G
	if int(255.99*color.Y) < 256 {
		img.Color.G = int(255.99 * color.Y)
	} else if int(255.99*color.Y) < 0 {
		img.Color.G = 0
	} else {
		img.Color.G = 255
	}

	// B
	if int(255.99*color.Z) < 256 {
		img.Color.B = int(255.99 * color.Z)
	} else if int(255.99*color.Z) < 0 {
		img.Color.B = 0
	} else {
		img.Color.B = 255
	}
}
