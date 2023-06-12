package indicators

//Calculate Relative Strength Index (RSI)
func RSI(input []float64, period int) []float64 {
	if len(input) < period {
		return nil
	}

	rsi := make([]float64, len(input))
	gains := make([]float64, len(input))
	losses := make([]float64, len(input))
	for i:=1; i<len(input); i++ {
		gains[i], losses[i] = calculateGainLoss(input[i], input[i-1])
	}

	avgGains,_ := MA(gains, period, WILDER)
	avgLosses,_ := MA(losses, period, WILDER)
	for i:=period; i<len(input); i++ {
		rsi[i] = 100.0 * avgGains[i] / ( avgGains[i] + avgLosses[i] )
	}

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