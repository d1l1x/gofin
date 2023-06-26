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

	if len(ind.input) < ind.period {
		return nil
	}

	rsi := make([]float64, len(ind.input))

	sumGains := 0.0
	sumLosses := 0.0

	for i := 1; i < ind.period; i++ {
		gain, loss := ind.calculateGainLoss(ind.input[i], ind.input[i-1])
		sumGains += gain
		sumLosses += loss
	}
	sumGains /= float64(ind.period)
	sumLosses /= float64(ind.period)

	rsi[ind.period-1] = 100.0 - (100.0 / (1.0 + sumGains/sumLosses))

	for i := ind.period; i < len(ind.input); i++ {
		gain, loss := ind.calculateGainLoss(ind.input[i], ind.input[i-1])
		sumGains = (sumGains*float64(ind.period-1) + gain) / float64(ind.period)
		sumLosses = (sumLosses*float64(ind.period-1) + loss) / float64(ind.period)
		rsi[i] = 100.0 - (100.0 / (1.0 + sumGains/sumLosses))
	}
	// first value is just to start the averaging
	rsi[ind.period-1] = 0.0
	ind.values = rsi
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
