package indicators

import (
	"math"
)

type AverageTrueRange struct {
	BarHistoryIndicator
}

func ATR(bars BarHistory, period int) *AverageTrueRange {
	return &AverageTrueRange{
		BarHistoryIndicator: NewBarHistoryIndicator(bars, period),
	}
}

type TrueRange struct {
	BarHistoryIndicator
}

func TR(bars BarHistory) *TrueRange {
	return &TrueRange{
		BarHistoryIndicator: NewBarHistoryIndicator(bars, 0),
	}
}

type TrueRangePercent struct {
	BarHistoryIndicator
}

func ATRP(bars BarHistory, period int) *TrueRangePercent {
	return &TrueRangePercent{
		BarHistoryIndicator: NewBarHistoryIndicator(bars, period),
	}
}

func (ind *AverageTrueRange) Compute() []float64 {
	if len(ind.input.Close) < ind.Period || len(ind.input.High) < ind.Period || len(ind.input.Low) < ind.Period {
		return nil
	}
	tr := TR(ind.input)
	trueRange := tr.Compute()
	atr, _ := MA(trueRange, ind.Period).Compute(WILDER)
	return atr
}

func (ind *TrueRangePercent) Compute() []float64 {
	atr := ATR(ind.input, ind.Period).Compute()
	res := make([]float64, len(atr))
	for i, val := range atr {
		res[i] = val / ind.input.Close[i] * 100
	}
	return res
}

func (ind *TrueRange) Compute() []float64 {
	tr := make([]float64, len(ind.input.Close))
	tr[0] = ind.input.High[0] - ind.input.Low[0]
	for i := 1; i < len(ind.input.Close); i++ {
		highLow := ind.input.High[i] - ind.input.Low[i]
		highClose := math.Abs(ind.input.High[i] - ind.input.Close[i-1])
		lowClose := math.Abs(ind.input.Low[i] - ind.input.Close[i-1])
		tr[i] = math.Max(highLow, math.Max(highClose, lowClose))
	}
	return tr
}
