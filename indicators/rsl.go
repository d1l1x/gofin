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
	err := CheckInput(ind.Input, ind.Period)
	if err != nil {
		return nil, err
	}
	res := make([]float64, len(ind.Input))
	for i := ind.Period; i < len(ind.Input); i++ {
		res[i] = ind.Input[i] / Mean(ind.Input[i-ind.Period:i+1])
	}
	return res, nil
}
