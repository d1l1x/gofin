package indicators

func ROC(input []float64, period int) *RateOfChange {
	return &RateOfChange{
		TimeSeriesIndicator: NewTimeSeriesIndicator(input, period),
	}
}

type RateOfChange struct {
	TimeSeriesIndicator
}

func (ind *RateOfChange) Compute() []float64 {
	roc := make([]float64, len(ind.input))

	for i := ind.Period; i < len(ind.input); i++ {
		roc[i] = ((ind.input[i] - ind.input[i-ind.Period]) / ind.input[i-ind.Period]) * 100
	}

	return roc
}
