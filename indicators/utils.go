package indicators

import (
	"fmt"
	"math"
)

func Mean(in []float64) float64 {
	res := 0.0
	for _, val := range in {
		res += val
	}
	return res / float64(len(in))
}

func StdDev(in []float64) float64 {
	mu := Mean(in)
	res := 0.0
	for _, val := range in {
		res += math.Pow(val-mu, 2)
	}
	return math.Sqrt(res / float64(len(in)))
}

func CheckInput(input []float64, period int) error {

	if input == nil {
		return fmt.Errorf("input is uninitialized: %v", input)
	}
	if period <= 0 {
		return fmt.Errorf("invalid period: %d", period)
	}
	if period >= len(input) {
		return fmt.Errorf("invalid period: %d >= %d", period, len(input))
	}
	return nil
}

func sliceAlmostEqual(a, b []float64, acc float64, args ...string) (bool, error) {
	if len(a) != len(b) {
		return false, fmt.Errorf("slices must have equal length: %d != %d", len(a), len(b))
	}
	msg := ""

	switch len(args) {
	case 1:
		msg = args[0]
	}

	for i := range a {
		diff := math.Abs(a[i] - b[i])
		if math.IsNaN(diff) {
			return false, fmt.Errorf("found NaN at index %d, %v, %v", i, a[i], b[i])
		}
		if diff >= acc {
			return false, fmt.Errorf("%s%v!=%v at index %d", msg, a[i], b[i], i)
		}
	}
	return true, nil
}
