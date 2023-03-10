package indicators

import "fmt"

type maType uint

const (
	SMA maType = iota
	EMA
	LWMA
)

func MA(input []float64, period int, matype maType) ([]float64, error) {
	 err := CheckInput(input, period)
	 if err != nil {
		return nil, err
	 }
	
	var weights []float64
	switch matype {
	case SMA: 
		weights = computeSwmaWeights(period)
		return wma(input, weights), nil
	case LWMA:
		weights = computeLwmaWeights(period)
		return wma(input, weights), nil
	case EMA:
		return ema(input, period), nil
	default:
		return nil, fmt.Errorf("moving average type not yet implemented.: %d", matype)
	}
}

func computeSwmaWeights(period int) []float64 {
	weights := make([]float64, period)
	b := float64(period)
	for i := range weights {
		weights[i] = 1.0 / b
	}
	return weights
}

func computeLwmaWeights(period int) []float64 {
	weights := make([]float64, period)
	b := float64(period)*(float64(period)+1.0)/2.0
	for i := range weights {
		weights[i] = (float64(period) - float64(i)) / b
	}
	return weights
}

func ema(input []float64, period int) []float64 {
	// res := make([]float64, len(input))
	// copy(res, input)

	weights := computeSwmaWeights(period)
	res := wma(input, weights)
	alpha := 2.0/(float64(period) + 1)
	// y_(i+1) = y_i + alpha * (x_(i+1) i y_i)
	for i:=period-1; i<len(input) - 1; i++ {
		res[i + 1] = res[i] + alpha * (input[i+1] - res[i]) 
	}

	return res
}

func wma(input []float64, weights []float64) []float64 {
	res := make([]float64, len(input))
	for i := range input {
		if i+1 >= len(weights) {
			for j, weight := range weights {
				res[i] += input[i-j] * weight
			}
		}
	}
	return res
}
