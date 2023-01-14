package indicators

import (
	"math"
)

func Mean(in []float64) float64 {
	res := 0.0
	for _, val := range in {
		res += val
	}
	return res/float64(len(in))
}

func StdDev(in []float64) float64 {
	mu := Mean(in)
	res := 0.0
	for _, val := range in {
		res += math.Pow(val - mu, 2)
	}
	return math.Sqrt(res/float64(len(in)))
}
