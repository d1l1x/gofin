package indicators

type RelativeStrengthLevy struct {
	TimeSeriesIndicator
}

func RSL(input []float64, period int) *RelativeStrengthLevy {
	return &RelativeStrengthLevy{
		TimeSeriesIndicator: NewTimeSeriesIndicator(input, period),
	}
}

func (ind *RelativeStrengthLevy) Compute() ([]float64, error) {
	err := CheckInput(ind.input, ind.period)
	if err != nil {
		return nil, err
	}
	res := make([]float64, len(ind.input))
	for i := ind.period; i < len(ind.input); i++ {
		res[i] = ind.input[i] / Mean(ind.input[i-ind.period:i+1])
	}
	return res, nil
}
