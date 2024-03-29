package indicators

import (
	"fmt"
)

type BBands struct {
	upper      []float64
	mean       []float64
	lower      []float64
	band_width []float64
}

type BollingerBands struct {
	TimeSeriesIndicator
}

func BB(input []float64, period int) *BollingerBands {
	return &BollingerBands{
		TimeSeriesIndicator: NewTimeSeriesIndicator(input, period),
	}
}

func (ind *BollingerBands) Compute(factor float64, method maType) (BBands, error) {
	err := CheckInput(ind.Input, ind.Period)
	if err != nil {
		return BBands{}, err
	}
	if factor < 0 {
		return BBands{}, fmt.Errorf("invalid factor: %v", factor)
	}
	res := BBands{mean: []float64{}, lower: []float64{}, upper: []float64{}, band_width: []float64{}}

	res.mean, err = MA(ind.Input, ind.Period).Compute(method)
	if err != nil {
		return BBands{}, err
	}

	res.upper = make([]float64, len(ind.Input))
	res.lower = make([]float64, len(ind.Input))
	res.band_width = make([]float64, len(ind.Input))
	stddev := 0.0
	for i := ind.Period - 1; i < len(ind.Input); i++ {
		stddev = factor * StdDev(ind.Input[i-ind.Period+1:i+1])
		res.upper[i] = res.mean[i] + stddev
		res.lower[i] = res.mean[i] - stddev
		res.band_width[i] = res.upper[i] - res.lower[i]
	}

	return res, nil
}
