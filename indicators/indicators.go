package indicators

type BarHistory struct {
	Open   []float64
	High   []float64
	Low    []float64
	Close  []float64
	Volume []int64
}

type Indicator interface {
	Compute() []float64
	SetInput(bars *BarHistory)
}

type GeneralIndicator struct {
	Values []float64
	Period int
}

type TimeSeriesIndicator struct {
	Input []float64
	GeneralIndicator
}

func (ind *TimeSeriesIndicator) SetInput(bars *BarHistory) {
	ind.Input = bars.Close
}

type BarHistoryIndicator struct {
	Input BarHistory
	GeneralIndicator
}

func (ind *BarHistoryIndicator) SetInput(bars *BarHistory) {
	ind.Input = *bars
}

func NewTimeSeriesIndicator(input []float64, period int) TimeSeriesIndicator {
	return TimeSeriesIndicator{
		Input: input,
		GeneralIndicator: GeneralIndicator{
			Values: []float64{},
			Period: period,
		},
	}
}

func NewBarHistoryIndicator(input BarHistory, period int) BarHistoryIndicator {
	return BarHistoryIndicator{
		Input: input,
		GeneralIndicator: GeneralIndicator{
			Values: []float64{},
			Period: period,
		},
	}
}
