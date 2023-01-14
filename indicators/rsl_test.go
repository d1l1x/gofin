package indicators

import "testing"


func TestRslPeriod1(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 1
	want := []float64{0,4./3,6./5,8./7,12./10,14./13,16./15,18./17,20./19}

	got, err := RSL(input, period)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRslPeriod2(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 2
	want := []float64{0,0,6./4,24./18,36./26,42./34,48./42,54./48,60./54}

	got, err := RSL(input, period)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}
	_, err = sliceAlmostEqual(got, want, 1e-9)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRslPeriod3(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 3
	want := []float64{0,0,0,32./20,48./30,56./40,64./50,72./60,80./68}

	got, err := RSL(input, period)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRslPeriod8(t *testing.T) {
	input := []float64{2,4,6,8,12,14,16,18,20}
	period := 8
	want := []float64{0,0,0,0,0,0,0,0,180./100}

	got, err := RSL(input, period)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}