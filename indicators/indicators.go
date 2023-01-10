package indicators

import (
	"fmt"
	"math"
)

func CheckInput(input []float64, period int) error {

	if input == nil {
		return fmt.Errorf("input is uninitialized: %v", input)
	}
	if period <= 0 {
		return fmt.Errorf("invalid period: %d", period)
	}
	if period > len(input) {
		return fmt.Errorf("invalid period: %d > %d", period, len(input))
	}
	return nil
}

func sliceAlmostEqual(a, b []float64, acc float64) (bool, error) {
	if len(a) != len(b) {
		return false, fmt.Errorf("slices must have equal length: %d != %d", len(a), len(b))
	}
	for i := range a{
		diff := math.Abs(a[i] - b[i])
		if diff >= acc {
			return false, fmt.Errorf("%v!=%v at index %d", a[i],b[i],i)
		}
	}
    return true, nil
}