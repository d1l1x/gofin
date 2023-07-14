package indicators

func RSI(input []float64, period int) *RelativeStrengthIndex {
	return &RelativeStrengthIndex{
		TimeSeriesIndicator: NewTimeSeriesIndicator(input, period),
	}
}

type RelativeStrengthIndex struct {
	TimeSeriesIndicator
}

func (ind *RelativeStrengthIndex) Compute() []float64 {

	if len(ind.Input) < ind.Period {
		return nil
	}

	rsi := make([]float64, len(ind.Input))

	sumGains := 0.0
	sumLosses := 0.0

	for i := 1; i < ind.Period; i++ {
		gain, loss := ind.calculateGainLoss(ind.Input[i], ind.Input[i-1])
		sumGains += gain
		sumLosses += loss
	}
	sumGains /= float64(ind.Period)
	sumLosses /= float64(ind.Period)

	rsi[ind.Period-1] = 100.0 - (100.0 / (1.0 + sumGains/sumLosses))

	for i := ind.Period; i < len(ind.Input); i++ {
		gain, loss := ind.calculateGainLoss(ind.Input[i], ind.Input[i-1])
		sumGains = (sumGains*float64(ind.Period-1) + gain) / float64(ind.Period)
		sumLosses = (sumLosses*float64(ind.Period-1) + loss) / float64(ind.Period)
		rsi[i] = 100.0 - (100.0 / (1.0 + sumGains/sumLosses))
	}
	// first value is just to start the averaging
	rsi[ind.Period-1] = 0.0
	ind.Values = rsi
	return rsi
}

func (ind *RelativeStrengthIndex) calculateGainLoss(currentPrice, previousPrice float64) (float64, float64) {
	gain := 0.0
	loss := 0.0
	if currentPrice > previousPrice {
		gain = currentPrice - previousPrice
	} else if currentPrice < previousPrice {
		loss = previousPrice - currentPrice
	} else {
		gain = 0.0
		loss = 0.0
	}

	return gain, loss
}
