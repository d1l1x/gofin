package indicators

import (
	"math"
)

type AverageDirectionalIndex struct {
	BarHistoryIndicator
}

func ADX(bars BarHistory, period int) *AverageDirectionalIndex {
	return &AverageDirectionalIndex{
		BarHistoryIndicator: NewBarHistoryIndicator(bars, period),
	}
}

// Calculate Average Directional Index (ADX)
func (ind *AverageDirectionalIndex) Compute() []float64 {
	if len(ind.Input.Close) < ind.Period {
		return nil
	}

	adip, adim := ind.DiPlusMinus()

	dx := 0.0
	for i := ind.Period; i < 2*ind.Period; i++ {
		dx += 100 * math.Abs(adip[i]-adim[i]) / (adip[i] + adim[i])
	}
	dx /= float64(ind.Period)

	adx := make([]float64, len(ind.Input.Close))
	adx[2*ind.Period-1] = dx
	for i := 2 * ind.Period; i < len(ind.Input.Close); i++ {
		adx[i] = (adx[i-1]*float64(ind.Period-1) + 100*math.Abs(adip[i]-adim[i])/(adip[i]+adim[i])) / float64(ind.Period)
	}

	return adx
}

func (ind *AverageDirectionalIndex) DiPlusMinus() ([]float64, []float64) {
	if len(ind.Input.Close) < ind.Period {
		return nil, nil
	}

	dp := make([]float64, len(ind.Input.Close))
	dm := make([]float64, len(ind.Input.Close))

	tr := TR(ind.Input).Compute()

	for i := 1; i < len(ind.Input.Close); i++ {

		upMove := ind.Input.High[i] - ind.Input.High[i-1]
		downMove := ind.Input.Low[i-1] - ind.Input.Low[i]

		if upMove > downMove && upMove > 0 {
			dp[i] = upMove
			dm[i] = 0
		}
		if downMove > upMove && downMove > 0 {
			dp[i] = 0
			dm[i] = downMove
		}
	}

	dip := 0.0
	dim := 0.0
	trp := 0.0

	for i := 0; i < ind.Period; i++ {
		dip += dp[i]
		dim += dm[i]
		trp += tr[i]
	}

	pdp := dip
	pdm := dim
	ptr := trp

	for i := ind.Period; i < len(ind.Input.Close); i++ {
		adp := pdp - pdp/float64(ind.Period) + dp[i]
		adm := pdm - pdm/float64(ind.Period) + dm[i]
		atr := ptr - ptr/float64(ind.Period) + tr[i]

		dp[i] = math.Round(100.0 * adp / atr)
		dm[i] = math.Round(100.0 * adm / atr)

		pdp = adp
		pdm = adm
		ptr = atr
	}

	return dp, dm

}
