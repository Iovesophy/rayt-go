package utils

func Mix(a float64, b float64, t float64) float64 {
	return a*(1-t) + b*t
}
