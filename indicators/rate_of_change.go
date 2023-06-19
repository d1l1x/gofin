package indicators

func ROC(input []float64, period int) *RateOfChange {
	return &RateOfChange{
		Indicator: NewIndicator(input,period),
	}
}

type RateOfChange struct {
	Indicator
}

func (ind *RateOfChange) Compute() []float64 {
	roc := make([]float64, len(ind.input))

	for i := ind.period; i < len(ind.input); i++ {
		roc[i] = ((ind.input[i] - ind.input[i-ind.period]) / ind.input[i-ind.period]) * 100
	}

	return roc
}