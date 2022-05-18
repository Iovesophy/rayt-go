package main

import (
	"fmt"
	"os"
	"strings"
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

func (img Image) CreateP3Data() Image {
	img.Header = img.CreateHeader()
	for j := 0; j < img.Y; j++ {
		for i := 0; i < img.X; i++ {
			r := float64(i) / float64(img.X)
			g := float64(j) / float64(img.Y)
			b := 0.2
			img.Color.R = int(255.99 * r)
			img.Color.G = int(255.99 * g)
			img.Color.B = int(255.99 * b)
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
	img.X = 200
	img.Y = 100
	img.MaxBright = 255
	result := img.CreateP3Data()

	filename := "test.ppm"
	img.CreateFile(filename, result.Header, result.Body.String())
}
