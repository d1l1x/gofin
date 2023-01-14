package indicators

import (
	"testing"
)

func TestBBBadFactor(t *testing.T) {
	input := []float64{90.7043,92.9001,92.9784,91.8021,92.6647,92.6843,92.3021,92.7725,92.5373,92.9490,93.2039,91.0669,89.8318,89.7435,90.3994,90.7387,88.0177,88.0867,88.8439,90.7781,90.5416,91.3894,90.6500}
	period := 5

	_, err := BollingerBands(input, period, -1.0, SMA)
	if err == nil {
		t.Errorf("Want error for bad factor: %v ", err)
	}

}

func TestBBPeriod5F0(t *testing.T) {
	input := []float64{90.7043,92.9001,92.9784,91.8021,92.6647,92.6843,92.3021,92.7725,92.5373,92.9490,93.2039,91.0669,89.8318,89.7435,90.3994,90.7387,88.0177,88.0867,88.8439,90.7781,90.5416,91.3894,90.6500}
	period := 5
	want := BB{mean:  []float64{0,0,0,0,92.209920,92.605920,92.486320,92.445140,92.592180,92.649040,92.752960,92.505920,91.917780,91.359020,90.849100,90.356060,89.746220,89.397200,89.217280,89.293020,89.253600,89.927940,90.440600},
	           upper:  []float64{0,0,0,0,92.209920,92.605920,92.486320,92.445140,92.592180,92.649040,92.752960,92.505920,91.917780,91.359020,90.849100,90.356060,89.746220,89.397200,89.217280,89.293020,89.253600,89.927940,90.440600},
	           lower:  []float64{0,0,0,0,92.209920,92.605920,92.486320,92.445140,92.592180,92.649040,92.752960,92.505920,91.917780,91.359020,90.849100,90.356060,89.746220,89.397200,89.217280,89.293020,89.253600,89.927940,90.440600},
			   band_width: []float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
			   }

	got, err := BollingerBands(input, period, 0.0, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got.mean, want.mean, 1e-4, "BB.mean: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.upper, want.upper, 1e-4, "BB.upper: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.lower, want.lower, 1e-4, "BB.lower: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.band_width, want.band_width, 1e-4, "BB.band_width: ")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBBPeriod5F2(t *testing.T) {
	input := []float64{90.7043,92.9001,92.9784,91.8021,92.6647,92.6843,92.3021,92.7725,92.5373,92.9490,93.2039,91.0669,89.8318,89.7435,90.3994,90.7387,88.0177,88.0867,88.8439,90.7781,90.5416,91.3894,90.6500}
	period := 5
	want := BB{mean:  []float64{0,0,0,0,92.209920,92.605920,92.486320,92.445140,92.592180,92.649040,92.752960,92.505920,91.917780,91.359020,90.849100,90.356060,89.746220,89.397200,89.217280,89.293020,89.253600,89.927940,90.440600},
			   upper: []float64{0,0,0,0,93.931999,93.445448,93.293910,93.164323,92.918883,93.086592,93.380300,94.009602,94.475378,94.320001,93.387131,91.377300,91.623830,91.685323,91.509657,91.755346,91.626784,92.426023,92.141806},
			   lower: []float64{0,0,0,0,90.487841,91.766392,91.678730,91.725957,92.265477,92.211488,92.125620,91.002238,89.360182,88.398039,88.311069,89.334820,87.868610,87.109077,86.924903,86.830694,86.880416,87.429857,88.739394},
			   band_width: []float64{0,0,0,0,3.444157,1.679055,1.615180,1.438365,0.653407,0.875104,1.254680,3.007364,5.115196,5.921961,5.076062,2.042480,3.755220,4.576245,4.584755,4.924652,4.746368,4.996166,3.402412},
			   }

	got, err := BollingerBands(input, period, 2.0, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got.mean, want.mean, 1e-4, "BB.mean: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.upper, want.upper, 1e-4, "BB.upper: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.lower, want.lower, 1e-4, "BB.lower: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.band_width, want.band_width, 1e-4, "BB.band_width: ")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBBPeriod20F2(t *testing.T) {
	input := []float64{90.7043,92.9001,92.9784,91.8021,92.6647,92.6843,92.3021,92.7725,92.5373,92.9490,93.2039,91.0669,89.8318,89.7435,90.3994,90.7387,88.0177,88.0867,88.8439,90.7781,90.5416,91.3894,90.6500}
	period := 20
	want := BB{mean:  []float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,91.250270,91.242135,91.166600,91.050180},
			   upper: []float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,94.534271,94.532306,94.369251,94.148503},
			   lower: []float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,87.966269,87.951964,87.963949,87.951857},
			   band_width: []float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,6.568003,6.580342,6.405301,6.196647},
			   }

	got, err := BollingerBands(input, period, 2.0, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got.mean, want.mean, 1e-4, "BB.mean: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.upper, want.upper, 1e-4, "BB.upper: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.lower, want.lower, 1e-4, "BB.lower: ")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sliceAlmostEqual(got.band_width, want.band_width, 1e-4, "BB.band_width: ")
	if err != nil {
		t.Fatal(err)
	}
}