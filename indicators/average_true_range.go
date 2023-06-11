package indicators

import (
	"math"
)

func TR(bars BarHistory) []float64 {
	tr  := make([]float64, len(bars.Close))
	tr[0] = bars.High[0] - bars.Low[0]
	for i:=1; i<len(bars.Close); i++ {
		highLow := bars.High[i] - bars.Low[i]
		highClose := math.Abs(bars.High[i] - bars.Close[i-1])
		lowClose := math.Abs(bars.Low[i] - bars.Close[i-1])
		tr[i] = math.Max(highLow, math.Max(highClose, lowClose))
	}
	return tr
}

func ATR(input BarHistory, period int) []float64 {
	if len(input.Close) < period || len(input.High) < period || len(input.Low) < period {
		return nil
	}
	tr := TR(input)
	atr,_ := MA(tr, period, WILDER) 
	return atr
}

func ATRP(input BarHistory, period int) []float64 {
	atr := ATR(input, period)
	res := make([]float64, len(atr))
	for i, val := range atr {
		res[i] = val / input.Close[i] * 100
	}
    return res
}