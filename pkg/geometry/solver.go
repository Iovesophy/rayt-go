package geometry

import "math"

func SphereNegSolver(a float64, b float64, c float64) float64 {
	return (-b - math.Sqrt(SphereDetector(a, b, c))) / (2.0 * a)
}

func SpherePosSolver(a float64, b float64, c float64) float64 {
	return (-b + math.Sqrt(SphereDetector(a, b, c))) / (2.0 * a)
}

func SphereDetector(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}
