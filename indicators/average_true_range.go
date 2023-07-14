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
	if len(ind.Input.Close) < ind.Period || len(ind.Input.High) < ind.Period || len(ind.Input.Low) < ind.Period {
		return nil
	}
	tr := TR(ind.Input)
	trueRange := tr.Compute()
	atr, _ := MA(trueRange, ind.Period).Compute(WILDER)
	return atr
}

func (ind *TrueRangePercent) Compute() []float64 {
	atr := ATR(ind.Input, ind.Period).Compute()
	res := make([]float64, len(atr))
	for i, val := range atr {
		res[i] = val / ind.Input.Close[i] * 100
	}
	return res
}

func (ind *TrueRange) Compute() []float64 {
	tr := make([]float64, len(ind.Input.Close))
	tr[0] = ind.Input.High[0] - ind.Input.Low[0]
	for i := 1; i < len(ind.Input.Close); i++ {
		highLow := ind.Input.High[i] - ind.Input.Low[i]
		highClose := math.Abs(ind.Input.High[i] - ind.Input.Close[i-1])
		lowClose := math.Abs(ind.Input.Low[i] - ind.Input.Close[i-1])
		tr[i] = math.Max(highLow, math.Max(highClose, lowClose))
	}
	return tr
}
