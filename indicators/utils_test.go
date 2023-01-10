package indicators

import (
	"testing"
)

func TestMeanPositive(t *testing.T) {
	in := []float64{1,2}
	want := 1.5
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