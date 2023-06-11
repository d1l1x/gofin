package indicators

//Calculate Relative Strength Index (RSI)
func RSI(input []float64, period int) []float64 {
	if len(input) < period {
		return nil
	}

	rsi := make([]float64, len(input))
	// var gainSum, lossSum float64
	// // for i:=period-1; i<len(input) - 1; i++ {
	// for i:=period; i<len(input) - 1; i++ {
	// 	// for j:=i-period+1; j<=i; j++ {
	// 	gain := 0.0
	// 	loss := 0.0
	// 	for j:=i-period; j<i; j++ {
	// 		gain, loss = calculateGainLoss(input[j+1], input[j])
	// 		gainSum += gain
	// 		lossSum += loss
	// 	}
	// 	// avgGain := gainSum / float64(period)
	// 	// avgLoss := lossSum / float64(period)
	// 	avgGain := (gainSum - gain) * float64(period - 1) + gain
	// 	avgLoss := (lossSum - loss) * float64(period - 1) + loss
	// 	rs := avgGain / avgLoss
	// 	rsi[i] = 100.0 - (100.0 / (1.0 + rs))
	// }
	// //Gains and Losses
	gains := make([]float64, len(input))
	losses := make([]float64, len(input))
	for i:=1; i<len(input); i++ {
		gains[i], losses[i] = calculateGainLoss(input[i], input[i-1])
	}

	avgGains,_ := MA(gains, period, WILDER)
	avgLosses,_ := MA(losses, period, WILDER)
	for i:=period; i<len(input) - 1; i++ {

		rs := avgGains[i] / avgLosses[i]
		rsi[i] = 100.0 - (100.0 / (1.0 + rs))
	}

	return rsi
}

// calculateGainLoss calculates the gain and loss for a given day.
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


// func RSI(data []float64, period int) []float64 {
// 	if len(data) < period {
//         return nil
//     }

//     gains := make([]float64, len(data))
//     losses := make([]float64, len(data))

//     for i := 1; i < len(data); i++ {
//         diff := data[i] - data[i-1]
//         if diff >= 0 {
//             gains[i] = diff
//         } else {
//             losses[i] = -diff
//         }
//     }

//     avgGain := calculateAverage(gains[:period])
//     avgLoss := calculateAverage(losses[:period])

//     rsi := make([]float64, len(data))
//     rsi[period-1] = 100 - (100 / (1 + avgGain/avgLoss))

//     for i := period; i < len(data); i++ {
//         avgGain = ((avgGain * float64(period-1)) + gains[i]) / float64(period)
//         avgLoss = ((avgLoss * float64(period-1)) + losses[i]) / float64(period)

//         rsi[i] = 100 - (100 / (1 + avgGain/avgLoss))
//     }

// 	fmt.Println(rsi)

//     return rsi
// }

// func calculateAverage(data []float64) float64 {
//     sum := 0.0
//     for _, value := range data {
//         sum += value
//     }
//     return sum / float64(len(data))
// }