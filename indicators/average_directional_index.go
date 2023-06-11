package indicators

import (
	"math"
)

func DiPlusMinus(input BarHistory, period int) ([]float64, []float64) {
	if len(input.Close) < period {
		return nil,nil
	}

	dp := make([]float64, len(input.Close))
	dm := make([]float64, len(input.Close))

	for i:=1; i<len(input.Close); i++ {

		upMove := input.High[i] - input.High[i-1]
		downMove := input.Low[i-1] - input.Low[i]

		if downMove > upMove && downMove > 0 {
			dp[i] = 0
			dm[i] = downMove
		} else if upMove > downMove && upMove > 0 {
			dp[i] = upMove
			dm[i] = 0
		} else {
			dp[i] = 0
			dm[i] = 0
		}
	}
	atr := ATR(input, period)

	adip, _ := MA(dp, period, WILDER)
	adim, _ := MA(dm, period, WILDER)

	for i:=period-1; i<len(input.Close); i++ {
		adip[i] = math.Round(100 * adip[i] / atr[i])
		adim[i] = math.Round(100 * adim[i] / atr[i])
	}

	return adip, adim

}

// Calculate Average Directional Index (ADX)
func ADX(input BarHistory, period int) []float64 {
	if len(input.Close) < period {
		return nil
	}

	adip,adim := DiPlusMinus(input, period)

	dx := make([]float64, len(adip))
	for i, val := range adip {
		dx[i] = math.Round(math.Abs(val - adim[i]) / math.Abs(val + adim[i]) * 100)
	}
	dx[0] = 0.0

	adx,_ := MA(dx, period, WILDER)

	return adx
}
