package indicators

import (
	"fmt"
)

type BB struct {
	upper []float64
	mean []float64
	lower []float64
	band_width []float64
}

func BollingerBands(in []float64, period int, factor float64, method maType) (BB,error) {
	 err := CheckInput(in, period)
	 if err != nil {
		return BB{}, err
	 }
	if factor < 0 {
		return BB{}, fmt.Errorf("invalid factor: %v", factor)
	}
	res := BB{mean: []float64{}, lower: []float64{}, upper: []float64{}, band_width: []float64{}}

	res.mean,err = MA(in,period,method)
	if err != nil {
		return BB{},err
	}

	res.upper = make([]float64, len(in))
	res.lower = make([]float64, len(in))
	res.band_width = make([]float64, len(in))
	stddev := 0.0
	for i := period - 1; i<len(in); i++ {
		stddev = factor*StdDev(in[i-period+1:i+1])
		res.upper[i] = res.mean[i] + stddev
		res.lower[i] = res.mean[i] - stddev
		res.band_width[i] = res.upper[i] - res.lower[i]
	}

	return res,nil
}