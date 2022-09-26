package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"

	"rayt-go/pkg/ray"
	"rayt-go/pkg/scene"

	"github.com/golang/geo/r3"
)

type RGB struct {
	R int
	G int
	B int
}

const (
	HeaderFormat = "%s\n%d %d\n%d\n"
	BodyFormat   = "%d %d %d\n"
)

type Image struct {
	Format    string
	X         int
	Y         int
	MaxBright int
	Header    string
	Body      strings.Builder
	Color     RGB
}

func (img Image) CreateHeader() string {
	return fmt.Sprintf(HeaderFormat, img.Format, img.X, img.Y, img.MaxBright)
}

func Gcd(m, n uint64) uint64 {
	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	a := new(big.Int).SetUint64(m)
	b := new(big.Int).SetUint64(n)
	return z.GCD(x, y, a, b).Uint64()
}

func (img Image) CreateP3Data() Image {
	gcd := Gcd(uint64(img.X), uint64(img.Y))
	x := float64(img.X) / float64(gcd)
	y := float64(img.Y) / float64(gcd)
	fmt.Println(x, y)
	img.Header = img.CreateHeader()
	lowerLeftCorner := r3.Vector{
		X: -x,
		Y: -y,
		Z: -1.0,
	}
	// lowerLeftCornerから基底ベクトルを求める
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
	for j := 0; j < img.Y; j++ {
		for i := 0; i < img.X; i++ {
			h := float64(i) / float64(img.X)
			v := float64(j) / float64(img.Y)
			r := ray.New(
				origin,
				lowerLeftCorner.Add(
					horizontal.Mul(h).Add(vertical.Mul(v)),
				),
			)
			pixel := scene.Pixel(r)
			img.Color.R = int(255.99 * pixel.X)
			img.Color.G = int(255.99 * pixel.Y)
			img.Color.B = int(255.99 * pixel.Z)
			img.Body.WriteString(fmt.Sprintf(BodyFormat, img.Color.R, img.Color.G, img.Color.B))
		}
	}
	return img
}

func (img Image) CreateFile(filename string, header string, body string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(header)
	if err != nil {
		return err
	}
	_, err = file.WriteString(body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	img := Image{}
	img.Format = "P3"
	img.X = 800
	img.Y = 400
	img.MaxBright = 255
	result := img.CreateP3Data()

	filename := "test.ppm"
	img.CreateFile(filename, result.Header, result.Body.String())
}
