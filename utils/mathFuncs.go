package utils

import "math"

func Float64Equals(x, y, accuracy float64) bool {
	return math.Abs(x-y) <= accuracy
}
