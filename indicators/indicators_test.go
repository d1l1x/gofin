package indicators

import (
	"testing"
)


func TestCheckInputZeroPeriod(t *testing.T) {
	input := []float64{1, 2, 3}
	period := 0

	err := CheckInput(input, period)
	if err == nil {
		t.Errorf("Want error for invalid period: %v ", err)
	}
}

func TestCheckInputNegativePeriod(t *testing.T) {
	input := []float64{1, 2, 3}
	period := -2

	err := CheckInput(input, period)
	if err == nil {
		t.Errorf("Want error for invalid period: %v ", err)
	}
}

func TestCheckInputPeriodToLarge(t *testing.T) {
	input := []float64{1, 2, 3}
	period := 10

	err := CheckInput(input, period) 
	if err == nil {
		t.Errorf("Want error for invalid period: %v ", err)
	}
}

func TestCheckInputUninitializedInput(t *testing.T) {
	var input []float64
	period := 0

	err := CheckInput(input, period)
	if err == nil {
		t.Errorf("Want error for uninitialized input: %v ", err)
	}
}

func TestSliceAlmostEqualUnequalLength(t *testing.T) {
	a := []float64{1,2}
	b := []float64{1,2,3}

	_, err := sliceAlmostEqual(a,b, 1e-9)
	if err == nil {
		t.Errorf("Want error for unequal length: %v ", err)
	}
}


func TestSliceAlmostEqualNotEqual(t *testing.T) {
	a := []float64{1,2,4}
	b := []float64{1,2,3}

	_, err := sliceAlmostEqual(a,b, 1e-9)
	if err == nil {
		t.Errorf("Want error for not equal: %v ", err)
	}
}

func TestSliceAlmostEqualIsEqual(t *testing.T) {
	a := []float64{1,2,3}
	b := []float64{1,2,3}

	_, err := sliceAlmostEqual(a,b, 1e-9)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
}