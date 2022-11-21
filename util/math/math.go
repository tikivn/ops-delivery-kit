package math

import (
	"math"

	"github.com/pkg/errors"
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

func Round(val float64, roundPoint float64, precision int) (newVal float64, err error) {
	if roundPoint >= 1 || roundPoint <= 0 {
		return 0, errors.New("Invalid round point (must greater than 0 and lower than 1)")
	}
	var round float64
	pow := math.Pow(10, float64(precision))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundPoint {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return newVal, nil
}
