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
	roc := make([]float64, len(ind.Input))

	for i := ind.Period; i < len(ind.Input); i++ {
		roc[i] = ((ind.Input[i] - ind.Input[i-ind.Period]) / ind.Input[i-ind.Period]) * 100
	}

	return roc
}
