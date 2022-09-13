package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"

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

type Ray struct {
	Origin    r3.Vector
	Direction r3.Vector
}

func NewRay(origin r3.Vector, direction r3.Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}

// 線形補完 パラメータtで補完している
func (r Ray) PointAtParameter(t float64) r3.Vector {
	return r3.Vector{
		X: r.Origin.X + t*r.Direction.X,
		Y: r.Origin.Y + t*r.Direction.Y,
		Z: r.Origin.Z + t*r.Direction.Z,
	}
}

func Solver(a float64, b float64, c float64) float64 {
	return (-b - math.Sqrt(Detector(a, b, c))) / (2.0 * a)
}

func Detector(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}

func HitSphere(center r3.Vector, radius float64, ray Ray) float64 {
	oc := ray.Origin.Sub(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - radius*radius
	if Detector(a, b, c) < 0 {
		return -1.0
	}
	return Solver(a, b, c)
}

func Color(ray Ray) r3.Vector {
	center := r3.Vector{
		X: 0.0,
		Y: 0.0,
		Z: -1.0,
	}
	unitVector := r3.Vector{
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}

	// Sphere
	sphereSize := 0.6
	t := HitSphere(center, sphereSize, ray)
	if t > 0.0 {
		n := ray.PointAtParameter(t).Sub(center).Normalize()
		result := r3.Vector{
			X: n.X,
			Y: n.Y,
			Z: n.Z,
		}.Add(
			unitVector,
		).Mul(
			0.5,
		)
		return result
	}

	// Background
	unit := ray.Direction.Normalize()
	t = 0.5*unit.Y + 1.0
	result := unitVector.Add(
		r3.Vector{
			X: 0.7,
			Y: 0.7,
			Z: 0.7,
		}.Mul(
			t,
		))
	return result
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
			ray := NewRay(
				origin,
				lowerLeftCorner.Add(
					horizontal.Mul(h).Add(vertical.Mul(v)),
				),
			)
			col := Color(ray)
			img.Color.R = int(255.99 * col.X)
			img.Color.G = int(255.99 * col.Y)
			img.Color.B = int(255.99 * col.Z)
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
