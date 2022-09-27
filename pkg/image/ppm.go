package image

import (
	"fmt"
	"math"
	"rayt-go/pkg/format"
	"rayt-go/pkg/geometry"
	"rayt-go/pkg/ray"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/utils"
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
	MaxBright int
	Header    string
	Body      strings.Builder
	Color     RGB
	World     []geometry.Hitable
}

func (img Elements) CreateHeader() string {
	return fmt.Sprintf(format.Header, img.Format, img.X, img.Y, img.MaxBright)
}

func (img Elements) CreateP3Data() Elements {
	gcd := utils.Gcd(uint64(img.X), uint64(img.Y))
	x := float64(img.X) / float64(gcd)
	y := float64(img.Y) / float64(gcd)
	img.Header = img.CreateHeader()
	lowerLeftCorner := r3.Vector{
		X: -x,
		Y: -y,
		Z: -1.0,
	}
	// lowerLeftCorner から基底ベクトルを求める
	horizontal := r3.Vector{
		X: math.Abs(lowerLeftCorner.X * 2.0),
		Y: 0.0,
		Z: 0.0,
	}
	vertical := r3.Vector{
		X: 0.0,
		Y: math.Abs(lowerLeftCorner.Y * 2.0),
		Z: 0.0,
	}
	origin := r3.Vector{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	p := make([]string, img.X*img.Y)
	wg.Add(img.X * img.Y)
	for j := 0; j < img.Y; j++ {
		for i := 0; i < img.X; i++ {
			go Render(i, j, origin, lowerLeftCorner, horizontal, vertical, wg, mu, &img, p)
		}
	}
	wg.Wait()
	for _, v := range p {
		img.Body.WriteString(v)
	}
	return img
}

func Render(i int, j int, origin r3.Vector, lowerLeftCorner r3.Vector, horizontal r3.Vector, vertical r3.Vector, wg *sync.WaitGroup, mu *sync.Mutex, img *Elements, p []string) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	h := float64(i) / float64(img.X)
	v := float64(j) / float64(img.Y)
	ray := ray.New(
		origin,
		lowerLeftCorner.Add(
			horizontal.Mul(h).Add(vertical.Mul(v)),
		),
	)
	sky := scene.Color{X: 0.5, Y: 0.7, Z: 1.0}
	world := geometry.New(img.World)
	color := sky.Pixel(ray, world)
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
