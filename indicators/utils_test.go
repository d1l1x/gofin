package indicators

import (
	"testing"
	"math"
)

func TestMeanPositive(t *testing.T) {
	in := []float64{1,2}
	want := 1.5
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}

func TestMeanPositive2(t *testing.T) {
	in := []float64{2,4,6,8,12,14,16,18,20}
	want := 100./9
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}

func TestMeanNegative(t *testing.T) {
	in := []float64{-1,-2}
	want := -1.5
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}

func TestMeanM1P1(t *testing.T) {
	in := []float64{-1,1}
	want := 0.0
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}

func TestMeanSingleZero(t *testing.T) {
	in := []float64{0}
	want := 0.0
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}
 
func TestMeanMultipleSameValues(t *testing.T) {
	in := []float64{2,2,2}
	want := 2.0
	got := Mean(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}


func TestStdDevSingleZero(t *testing.T) {
	in := []float64{0}
	want := 0.0
	got := StdDev(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}

}

func TestStdDevSameValues(t *testing.T) {
	in := []float64{1,1,1}
	want := 0.0
	got := StdDev(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}
}

func TestStdDevSimple(t *testing.T) {
	in := []float64{2,4}
	want := 1.0
	got := StdDev(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}
}

func TestStdDevPositive(t *testing.T) {
	in := []float64{2,4,6,8,12,14,16,18,20}
	// compute with WolframAlpha
	want := 4.0*math.Sqrt(185.0)/9.0
	got := StdDev(in)
	_, err := sliceAlmostEqual([]float64{got}, []float64{want},1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStdDevNegativePositive(t *testing.T) {
	in := []float64{-2,-4,-6,-8,-12,14,16,18,20}
	// compute with WolframAlpha
	want := 12.0
	got := StdDev(in)
	if got != want {
		t.Fatalf("got != want: %v != %v ", got, want)
	}
}