package indicators

import (
	"testing"
)

var TestBars = BarHistory{
	Open: []float64{127.25, 127, 126.47, 127.75, 130, 130.4, 129.86, 130.86, 130.3, 132.1, 133.75, 134.41, 133.44, 133.44, 134.87, 136.14, 136.51, 137.91, 140.09, 143.5, 141.6, 142.76, 146.17, 144, 148},
	High: []float64{127.75, 128.19, 127.44, 130.49, 130.6, 130.89, 132.55, 131.51, 132.41, 134.08, 134.32, 134.64, 133.89, 135.245, 136.49, 137.41, 137.33, 140, 143.15, 144.89, 144.06, 145.65, 146.32, 147.46, 149.57},
	Low: []float64{126.52, 125.94, 126.1, 127.07, 129.39, 128.461, 129.65, 130.24, 129.21, 131.62, 133.23, 132.93, 132.81, 133.35, 134.35, 135.87, 135.76, 137.745, 140.07, 142.66, 140.665, 142.6522, 144, 143.63, 147.68},
	Close: []float64{127.13, 126.11, 127.35, 130.48, 129.64, 130.15, 131.79, 130.46, 132.3, 133.98, 133.7, 133.41, 133.11, 134.78, 136.33, 136.96, 137.27, 139.96, 142.02, 144.57, 143.24, 145.11, 144.5, 145.64, 149.15},
	Volume: []int64{55542559, 70604343, 52839509, 94601581, 62273470, 91339136, 96268773, 108467705, 78643119, 74473566, 59921548, 68358848, 69429845, 61910717, 62012618, 62951963, 52094428, 78690866, 107697800, 104465525, 105315267, 99223602, 76056394, 100543496, 126131926},
}


// ADX: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 39.8, 41.44, 43.752000000000000, 44.401600000000000, 48.72128000000000, 53.577024000000000, 58.6616192, 61.92929536000000, 66.743436288, 71.5947490304, 76.07579922432, 70.460639379456, 67.7685115035648, 66.21480920285180, 66.37184736228150, 68.29747788982520},
// ATRP: []float64{0, 0, 0, 0, 1.4578833693304600, 1.5349980791394500, 1.6528112906897300, 1.573349685727420, 1.724923356009070, 1.7298543663233300, 1.5498332864622300, 1.4989143866276900, 1.3641058872811900, 1.394575662003260, 1.4169201656263500, 1.353205176284110, 1.308865627996630, 1.4170790783372200, 1.5664519803074400, 1.6280971722667900, 1.8598101546752800, 1.881850770067880, 1.8329431985857400, 1.9808310474904700, 2.074358612170370},



// Moving average
func TestSma(t *testing.T) {
	period := 5
	want := []float64{0, 0, 0, 0, 128.142, 128.746, 129.882, 130.504, 130.868, 131.736, 132.446, 132.77, 133.3, 133.796, 134.266, 134.918, 135.69, 137.06, 138.508, 140.156, 141.412, 142.98, 143.888, 144.612, 145.528}


	got, err := MA(TestBars.Close, period, SMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEma(t *testing.T) {
	period := 5
	want := []float64{127.13, 126.79000000000000, 126.97666666666700, 128.14444444444400, 128.64296296296300, 129.1453086419750, 130.02687242798400, 130.17124828532200, 130.8808321902150, 131.91388812681000, 132.5092587512070, 132.80950583413800, 132.9096705560920, 133.5331137040610, 134.46540913604100, 135.29693942402700, 135.9546262826850, 137.2897508551230, 138.86650057008200, 140.7676670467210, 141.5917780311480, 142.76451868743200, 143.34301245828800, 144.1086749721920, 145.78911664812800}


	got, err := MA(TestBars.Close, period, EMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

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


func TestLwma(t *testing.T) {
	period := 5
	want := []float64{0, 0, 0, 0, 128.76800000000000, 129.43733333333300, 130.452, 130.64466666666700, 131.24333333333300, 132.28066666666700, 132.93533333333300, 133.25666666666700, 133.37000000000000, 133.86333333333300, 134.708, 135.60600000000000, 136.39000000000000, 137.81333333333300, 139.46666666666700, 141.48733333333300, 142.51533333333300, 143.74800000000000, 144.25466666666700, 144.83866666666700, 146.35133333333300}


	got, err := MA(TestBars.Close, period, LWMA)
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v ", err)
	}

	_, err = sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

// RSI
func TestRSI(t *testing.T) {
    period := 5
	want := []float64{0, 0, 0, 0, 0, 73.64005458396610, 80.00012052113340, 64.27840280298390, 73.33908033718810, 79.32442189375520, 75.78014216551850, 71.63629101698800, 66.90539767215500, 77.32518815773540, 83.39183496367500, 85.37926999702190, 86.38163827094260, 92.18966430904170, 94.45388040259210, 96.17131656968280, 80.01679812716840, 84.57159864721220, 77.37991196102190, 81.12871570629630, 88.47793300360870}


    got := RSI(TestBars.Close, period)

	_, err := sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

// ROC
func TestROC(t *testing.T) {
	period := 5
	want := []float64{0, 0, 0, 0, 0, 2.3755211201132800, 4.504004440567760, 2.4420887318413900, 1.394849785407740, 3.347732181425470, 2.72762197464462, 1.2292283177782900, 2.031273953702280, 1.874527588813310, 1.7539931333034800, 2.438294689603610, 2.8933363316093300, 5.146119750582220, 5.371716871939470, 6.044157558864510, 4.585280373831770, 5.711371749107610, 3.2437839382680800, 2.54893676946908, 3.1680154942242500}


    got := ROC(TestBars.Close, period)

	_, err := sliceAlmostEqual(got, want,1e-6)
	if err != nil {
		t.Fatal(err)
	}
}

// ATR
func TestATR(t *testing.T) {
	period := 5
	want := []float64{0, 0, 0, 0, 1.8900000000000100, 1.9978000000000000, 2.178240000000000, 2.0525920000000000, 2.282073600000000, 2.3176588800000000, 2.0721271040000000, 1.9997016832000000, 1.8157613465599900, 1.8796090772479900, 1.931687261798400, 1.853349809438720, 1.7966798475509800, 1.9833438780407800, 2.2246751024326200, 2.3537400819460900, 2.663992065556880, 2.7307536524455000, 2.6486029219564000, 2.8848823375651200, 3.0939058700521000}


	got := ATR(TestBars, period)

	_, err := sliceAlmostEqual(got, want,1e-6)
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