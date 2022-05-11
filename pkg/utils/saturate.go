package utils

func Saturate(x float64) float64 {
	if x < 0 {
		return 0
	} else if x > 1 {
		return 1
	}
	return x
}
