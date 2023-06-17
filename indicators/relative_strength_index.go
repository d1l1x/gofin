package indicators

//Calculate Relative Strength Index (RSI)
func RSI(input []float64, period int) []float64 {
	if len(input) < period {
		return nil
	}

	rsi := make([]float64, len(input))

	sumGains := 0.0
	sumLosses := 0.0

	for i:=1; i<period; i++ {
		gain, loss := calculateGainLoss(input[i], input[i-1])
		sumGains += gain
		sumLosses += loss
	}
	sumGains /= float64(period)
	sumLosses /= float64(period)

	rsi[period-1] = 100.0 - ( 100.0 / ( 1.0 + sumGains / sumLosses ) )

	for i:=period; i<len(input); i++ {
		gain, loss := calculateGainLoss(input[i], input[i-1])
		sumGains = (sumGains * float64(period-1)  + gain) / float64(period)
		sumLosses = (sumLosses * float64(period-1)  + loss) / float64(period)
		rsi[i] = 100.0 - ( 100.0 / ( 1.0 + sumGains / sumLosses ) )
	}
	// first value is just to start the averaging
	rsi[period-1] = 0.0
	return rsi
}

func calculateGainLoss(currentPrice, previousPrice float64) (float64, float64) {
	gain := 0.0
	loss := 0.0
	if currentPrice > previousPrice {
		gain = currentPrice - previousPrice
	} else if currentPrice < previousPrice{
		loss = previousPrice - currentPrice
	} else {
		gain = 0.0
		loss = 0.0
	}

	return gain, loss
}