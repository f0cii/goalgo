package math2

import (
	"math"
	"strconv"
)

func ToFloat64(x string) float64 {
	v, _ := strconv.ParseFloat(x, 64)
	return v
}

// Round returns the nearest integer, rounding ties away from zero.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

// RoundToEven returns the nearest integer, rounding ties to an even number.
func RoundToEven(x float64) float64 {
	t := math.Trunc(x)
	odd := math.Remainder(t, 2) != 0
	if d := math.Abs(x - t); d > 0.5 || (d == 0.5 && odd) {
		return t + math.Copysign(1, x)
	}
	return t
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
