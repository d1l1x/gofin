/*
Tests for  SMA indicator
*/
package indicators

import (
	"testing"
)


// SMA
func TestSmaPeriod1(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 1
	want := []float64{2,4,6,8,12,14,16,18,20}

	got, err := MA(input, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSmaPeriod2(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 2
	want := []float64{0,3,5,7,10,13,15,17,19}

	got, err := MA(input, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSmaPeriod3(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 3
	want := []float64{0,0,4,6,8.666666667,11.333333333,14,16,18}

	got, err := MA(input, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSmaPeriodSameAsInputSize(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 9
	want := []float64{0,0,0,0,0,0,0,0,11.11111111}

	got, err := MA(input, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

// WMA
func TestComputeLwmaWeights1(t *testing.T) {
	period := 1
	want := []float64{1.0}
	got := computeLwmaWeights(period)

	_, err := sliceAlmostEqual(got, want,1e-6)
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

	_, err := sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLwmaPeriod1(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 1
	want := []float64{2,4,6,8,12,14,16,18,20}

	got, err := MA(input, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLwmaPeriod2(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 2
	want := []float64{0,10.0/3.0,16.0/3.0,22.0/3.0,32.0/3.0,40.0/3.0,46.0/3.0,52.0/3.0,58.0/3.0}

	got, err := MA(input, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLwmaPeriod3(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 3
	want := []float64{0,0,28.0/6.0,40.0/6.0,58.0/6.0,74.0/6.0,88.0/6.0,100.0/6.0,112.0/6.0}

	got, err := MA(input, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLwmaPeriod5(t *testing.T) {
	input := []float64{22.73,22.71,22.57,22.59,22.72}
	period := 5
	want := []float64{0,0,0,0,22.65466666667}

	got, err := MA(input, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmaPeriod5(t *testing.T) {
	input := []float64{1,2,3,4,5,6,7,8,9,10}
	period := 5
	want := []float64{0,0,0,0,3,4,5,6,7,8}

	got, err := MA(input, period, EMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}


func TestEmaPeriod14(t *testing.T) {
	input := []float64{7509.7, 7316.17, 7251.52, 7195.79, 7188.3, 7246, 7296.24, 7385.54, 7220.24, 7168.36, 7178.68, 6950.56, 7338.91, 7344.48, 7356.7, 7762.74, 8159.01, 8044.44, 7806.78, 8200, 8016.22, 8180.76, 8105.01, 8813.04, 8809.17, 8710.15, 8892.63, 8908.53, 8696.6, 8625.17, 8717.89, 8655.93, 8378.44, 8422.13, 8329.5, 8590.48, 8894.54, 9400, 9289.18, 9500, 9327.85, 9377.17, 9329.39, 9288.09, 9159.37, 9618.42, 9754.63, 9803.42, 9902, 10173.97, 9850.01, 10268.98, 10348.78, 10228.67, 10364.04, 9899.78, 9912.89, 9697.15, 10185.17, 9595.72, 9612.76, 9696.13, 9668.13, 9965.21, 9652.58, 9305.4, 8779.36, 8816.5, 8703.84, 8527.74, 8528.95, 8917.34, 8755.45, 8753.28, 9066.65, 9153.79, 8893.93, 8033.7, 7936.25, 7885.92, 7934.57, 4841.67, 5622.74, 5169.37, 5343.64, 5033.42, 5324.99, 5406.92, 6181.18, 6210.14, 6187.78, 5813.15, 6493.51, 6768.64, 6692.22, 6760.72, 6376.03, 6253.08, 5870.9, 5947.01}
	period := 14
	want := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7256.463571428571, 7269.828428571429, 7335.549971428572, 7445.344641904762, 7525.224022984127, 7562.764819919577, 7647.729510596967, 7696.861575850705, 7761.3813657372775, 7807.1985169723075, 7941.310714709333, 8057.025286081423, 8144.108581270566, 8243.911437101156, 8332.527245487669, 8381.070279422645, 8413.616908832959, 8454.186654321898, 8481.08576707898, 8467.399664801782, 8461.363709494877, 8443.781881562227, 8463.341630687262, 8520.834746595627, 8638.056780382876, 8724.87320966516, 8828.223448376471, 8894.840321926275, 8959.150945669438, 9008.516152913513, 9045.792665858378, 9060.936310410594, 9135.267469022516, 9217.849139819513, 9295.925254510244, 9376.735220575545, 9483.033191165472, 9531.96343234341, 9630.232308030954, 9726.038666960161, 9793.05617803214, 9869.187354294521, 9873.266373721919, 9878.54952389233, 9854.362920706686, 9898.470531279128, 9858.103793775244, 9825.391287938544, 9808.15644954674, 9789.48625627384, 9812.91608877066, 9791.537943601239, 9726.719551121074, 9600.404944304932, 9495.884285064274, 9390.278380389038, 9275.273263003834, 9175.763494603323, 9141.307028656214, 9089.859424835386, 9044.982168190667, 9047.871212431912, 9061.99371744099, 9039.58522178219, 8905.467192211232, 8776.238233249735, 8657.529135483102, 8561.134584085356, 8065.205972873975, 7739.543843157445, 7396.853997403119, 7123.092131082703, 6844.469180271676, 6641.871956235453, 6477.211695404059, 6437.740802683518, 6407.394028992382, 6378.112158460064, 6302.783870665389, 6328.214021243338, 6386.937485077559, 6427.641820400551, 6472.052244347145, 6459.249278434192, 6431.760041309633, 6356.978702468349, 6302.3162088059025}

	got, err := MA(input, period, EMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}