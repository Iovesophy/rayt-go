package utils

func Clamp(x float64, a float64, b float64) float64 {
	if x < a {
		return a
	} else if x > b {
		return b
	}
	return x
}
