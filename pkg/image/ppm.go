package image

import (
	"fmt"
	"math/rand"
	"rayt-go/pkg/format"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/scene/camera"
	"strings"
	"sync"
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
	World     []geometry.Hitable
	Camera    camera.Parts
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
	sky := scene.Color{X: 0.5, Y: 0.7, Z: 1.0}
	world := geometry.New(img.World)
	color := scene.NewVector(0, 0, 0)
	for s := 0; s < img.Sampling; s++ {
		h := (float64(i) + rand.Float64()) / float64(img.X)
		v := (float64(j) + rand.Float64()) / float64(img.Y)
		ray := camera.GetCameraRay(h, v)
		color = color.Add(sky.Pixel(ray, world))
	}
	color = color.Mul(1.0 / float64(img.Sampling))
	if int(255.99*color.X) < 256 {
		img.Color.R = int(255.99 * color.X)
	} else {
		img.Color.R = 255
	}
	if int(255.99*color.Y) < 256 {
		img.Color.G = int(255.99 * color.Y)
	} else {
		img.Color.G = 255
	}
	if int(255.99*color.Z) < 256 {
		img.Color.B = int(255.99 * color.Z)
	} else {
		img.Color.B = 255
	}
	p[i+j*img.X] = fmt.Sprintf(format.Body, img.Color.R, img.Color.G, img.Color.B)
}
