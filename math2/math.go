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

// Round4BitMEX 类似四舍五入法
// XBT: precision=0 0,0.5,1.0,1.5...
// ETH: precision=1 0.05,0.10,0.15...
func Round4BitMEX(x float64, precision int) float64 {
	if precision == 0 {
		return round4BitMEX(x)
	}
	p := math.Pow(10, float64(precision))
	y := float64(round4BitMEX(x*p)) / p
	return y
}

func round4BitMEX(x float64) float64 {
	t := math.Trunc(x)
	if x > t+0.5 {
		t += 0.5
	}
	if d := math.Abs(x - t); d > 0.25 {
		return t + math.Copysign(0.5, x)
	}
	return t
}

// RoundToEven5 类似四舍五入法，规整到 0,0.5,1.0,1.5...
func RoundToEven5(x float64) float64 {
	t := math.Trunc(x)
	if x > t+0.5 {
		t += 0.5
	}
	if d := math.Abs(x - t); d > 0.25 {
		return t + math.Copysign(0.5, x)
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
