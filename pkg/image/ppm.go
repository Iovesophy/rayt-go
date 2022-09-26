package image

import (
	"fmt"
	"math"
	"rayt-go/pkg/ray"
	"rayt-go/pkg/scene"
	"rayt-go/pkg/utils"
	"strings"
	"sync"

	"github.com/golang/geo/r3"
)

const (
	HeaderFormat = "%s\n%d %d\n%d\n"
	BodyFormat   = "%d %d %d\n"
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
}

func (img Elements) CreateHeader() string {
	return fmt.Sprintf(HeaderFormat, img.Format, img.X, img.Y, img.MaxBright)
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
	color := scene.Pixel(ray)
	img.Color.R = int(255.99 * color.X)
	img.Color.G = int(255.99 * color.Y)
	img.Color.B = int(255.99 * color.Z)
	p[i+j*img.X] = fmt.Sprintf(BodyFormat, img.Color.R, img.Color.G, img.Color.B)
}
