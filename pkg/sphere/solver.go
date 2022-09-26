package sphere

import "math"

func Solver(a float64, b float64, c float64) float64 {
	return (-b - math.Sqrt(Detector(a, b, c))) / (2.0 * a)
}

func Detector(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}
