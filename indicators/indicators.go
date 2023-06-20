package indicators

type BarHistory struct {
	Open   []float64
	High   []float64
	Low    []float64
	Close  []float64
	Volume []int64
}

type GeneralIndicator struct {
	values []float64
	period int
}

type TimeSeriesIndicator struct {
	input []float64
	GeneralIndicator
}

type BarHistoryIndicator struct {
	input BarHistory
	GeneralIndicator
}

func NewTimeSeriesIndicator(input []float64, period int) TimeSeriesIndicator {
	return TimeSeriesIndicator{
		input: input,	
		GeneralIndicator: GeneralIndicator{
			values: []float64{}, 
			period: period,
		},
	}
}

func NewBarHistoryIndicator(input BarHistory, period int) BarHistoryIndicator {
	return BarHistoryIndicator{
		input: input,	
		GeneralIndicator: GeneralIndicator{
			values: []float64{}, 
			period: period,
		},
	}
}

func (i *GeneralIndicator) Last() float64 {
	return i.values[len(i.values)-1]
}