package utils

import "math"

func Radians(deg float64) float64 {
	return (deg / 180) * math.Pi
}
