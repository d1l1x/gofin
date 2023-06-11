package indicators

import (
	"testing"
	"fmt"
)

var ACC = 1e-1

var TestBars = BarHistory{
	Open : []float64{127.75, 130.00, 130.40, 129.86, 130.86, 130.30, 132.10, 133.75, 134.41, 133.44, 133.44, 134.87, 136.14, 136.51, 137.91, 140.09, 143.50, 141.60, 142.76, 146.17, 144.00, 148.00, 149.17, 148.53, 143.78, 143.51, 145.51, 145.89, 147.50, 148.20, 149.03, 144.84},
	High : []float64{130.49, 130.60, 130.89, 132.55, 131.51, 132.41, 134.08, 134.32, 134.64, 133.89, 135.245, 136.49, 137.41, 137.33, 140.00, 143.15, 144.89, 144.06, 145.65, 146.32, 147.46, 149.57, 150.00, 149.76, 144.07, 147.10, 146.13, 148.195, 148.72, 149.83, 149.21, 146.97},
	Low  : []float64{127.07, 129.39, 128.461, 129.65, 130.24, 129.21, 131.62, 133.23, 132.93, 132.81, 133.35, 134.35, 135.87, 135.76, 137.745, 140.07, 142.66, 140.67, 142.6522, 144.00, 143.63, 147.68, 147.09, 145.88, 141.67, 142.96, 144.63, 145.81, 146.92, 147.70, 145.55, 142.54},
	Close : []float64{130.48, 129.64, 130.15, 131.79, 130.46, 132.30, 133.98, 133.70, 133.41, 133.11, 134.78, 136.33, 136.96, 137.27, 139.96, 142.02, 144.57, 143.24, 145.11, 144.50, 145.64, 149.15, 148.48, 146.39, 142.45, 146.15, 145.40, 146.80, 148.56, 148.99, 146.77, 144.98},
	Volume : []int64{94601581, 62273470, 91339136, 96268773, 108467705, 78643119, 74473566, 59921548, 68358848, 69429845, 61910717, 62012618, 62951963, 52094428, 78690866, 107697800, 104465525, 105315267, 99223602, 76056394, 100543496, 126131926, 103957709, 88965258, 116689470, 95243988, 71526938, 75328510, 70914977, 71859963, 103551959, 114096651},
}

// Moving average
func TestSma(t *testing.T) {
	period := 5
	want := []float64{0, 0, 0, 0, 130.504, 130.868, 131.736, 132.446, 132.77, 133.3, 133.796, 134.266, 134.918, 135.69, 137.06, 138.508, 140.156, 141.412, 142.98, 143.888, 144.612, 145.528, 146.576, 146.832, 146.422, 146.524, 145.774, 145.438, 145.872, 147.18, 147.304, 147.22}

	got, err := MA(TestBars.Close, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEma(t *testing.T) {
	period := 5
	want := []float64{130.48,130.2,130.1833,130.7188,130.6325,131.1883,132.1189,132.6459,132.9006,132.9704,133.5736,134.4924,135.3149,135.9666,137.2977,138.8718,140.7712,141.5941,142.7660,143.3440,144.1093,145.7895,146.6863,146.5875,145.2083,145.5222,145.4815,145.9210,146.8006,147.5304,147.2769,146.5113}


	got, err := MA(TestBars.Close, period, EMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestComputeLwmaWeights1(t *testing.T) {
	period := 1
	want := []float64{1.0}
	got := computeLwmaWeights(period)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}
func TestComputeLwmaWeights5(t *testing.T) {
	period := 5
	want := []float64{5.0/15.0, 
		              4.0/15.0,
					  3.0/15.0,
					  2.0/15.0,
					  1.0/15.0}
	got := computeLwmaWeights(period)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}


func TestLwma(t *testing.T) {
	period := 5
	want := []float64{0.0,0.0,0.0,0.0,130.6446,131.2433,132.2806,132.9353,133.2566,133.3700,133.8633,134.708 ,135.6060,136.3900,137.8133,139.4666,141.4873,142.5153,143.7480,144.2546,144.8386,146.3513,147.3353,147.2733,145.8126,145.722 ,145.3473,145.6893,146.73,147.7693,147.6326,146.8580}


	got, err := MA(TestBars.Close, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWilderMA(t *testing.T) {
	period := 5
	want := []float64{0.0,0.0,0.0,0.0,130.504 ,130.8632,131.4865,131.9292,132.2253,132.4023,132.8778,133.5682,134.2466,134.8513,135.8730,137.1024,138.5959,139.5247,140.6418,141.4134,142.2587,143.6370,144.6056,144.9624,144.4599,144.7979,144.9183,145.2947,145.9477,146.5562,146.5989,146.2751}


	got, err := MA(TestBars.Close, period, WILDER)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

//RSI
// func TestRSI(t *testing.T) {
//     period := 5
// 	want := []float64{0.0,0.0,0.0,0.0,0.0,67.7811,75.5784,71.9507,67.7412,62.9767,75.1422,82.0030,84.2163,85.3263,91.6758,94.1139,95.9496,79.7371,84.3767,77.1716,80.9681,88.3955,80.8655,60.7024,38.2362,56.9425,52.8841,59.6022,66.9970,68.7443,51.2379,40.7728}

//     got := RSI(TestBars.Close, period)

// 	for i := 0; i < len(got); i++ {
// 		fmt.Println(i, got[i], want[i])
// 	}

// 	_, err := sliceAlmostEqual(got, want,1e-6)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// ROC
func TestROC(t *testing.T) {
	period := 5
	want := []float64{0.0,0.0,0.0,0.0,0.0,1.395,3.348,2.728,1.229,2.031,1.8745,1.7539,2.4,2.893,5.146,5.372,6.04,4.59,5.711,3.244,2.549,3.1680,3.658,0.882,-1.419,0.350 ,-2.514,-1.131,1.482 ,4.591 ,0.424 ,-0.288}


    got := ROC(TestBars.Close, period)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

// True Range
func TestTR(t *testing.T) {
	want := []float64{3.4200, 1.2100, 2.4290, 2.9000, 1.5500, 3.2000, 2.4600, 1.0900, 1.7100, 1.1000, 2.1300, 2.1400, 1.5400, 1.5700, 2.7300, 3.1900, 2.8700, 3.9050, 2.9978, 2.3200, 3.8300, 3.9300, 2.9100, 3.8800, 4.7200, 4.6497, 1.5200, 2.7950, 1.9177, 2.1300, 3.6600, 4.4300}

	got := TR(TestBars)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestATR(t *testing.T) {
	period := 5
	want := []float64{0.0,0.0,0.0,0.0,2.3018,2.4814,2.4771,2.1997,2.1017,1.8974,1.9449,1.9839,1.8951,1.8301,2.0101,2.2460,2.3708,2.6776,2.7417,2.6573,2.8918,3.0995,3.0616,3.2252,3.5242,3.7493,3.3034,3.2017,2.9449,2.7819,2.9575,3.2520}

	got := ATR(TestBars, period)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestATRP(t *testing.T) {
	period := 5
	want := []float64{0.0,0.0,0.0,0.0,1.7643,1.8756,1.8488,1.6452,1.5754,1.4254,1.4430,1.4552,1.3837,1.3332,1.4361,1.5815,1.6399,1.8693,1.8894,1.8390,1.9856,2.0781,2.0619,2.2032,2.4740,2.5653,2.2719,2.1810,1.9823,1.8672,2.0151,2.2431}

	got := ATRP(TestBars, period)

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestADX(t *testing.T) {
	period := 2
	want := []float64{0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,32,37.6,44.68,51.5440,56.0352,62.0281,67.8225,73.0580,67.4464,65.3571,64.2857,64.8285,67.0628,65.4502,58.7602,52.2081,43.5665,36.8532,35.6825,35.5460,37.6368,31.3094,30.6475}

	got := ADX(TestBars, period)

	for i := 0; i < len(got); i++ {
		fmt.Println(i, got[i], want[i])
	}

	_, err := sliceAlmostEqual(got, want,ACC)
	if err != nil {
		t.Fatal(err)
	}
}

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