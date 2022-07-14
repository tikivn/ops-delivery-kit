package util

import (
	"math"
)

const float64EqualityThreshold = 1e-9

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func Abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func MaxFloat64(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func MinFloat64(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

func Round2Nearest(x float64, y float64) float64 {
	if y <= 0 {
		return x
	}

	return math.Round(x*math.Pow(10, y)) / math.Pow(10, y)
}
