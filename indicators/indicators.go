package indicators

type BarHistory struct {
	Open   []float64
	High   []float64
	Low    []float64
	Close  []float64
	Volume []int64
}

type Indicator struct {
	input []float64
	values []float64
	period int
}

func (i *Indicator) Last() float64 {
	return i.values[len(i.values)-1]
}

func NewIndicator(input []float64, period int) Indicator {
	return Indicator{
		input: input, // or some other initial values
		values: []float64{}, // or some other initial values
		period: period,
	}
}