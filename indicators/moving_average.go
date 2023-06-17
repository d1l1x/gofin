package indicators

import "fmt"

type maType uint

const (
	SMA maType = iota
	EMA
	LWMA
	WILDER
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
	case WILDER:
		return wilder(input, period), nil
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
	weights := computeSwmaWeights(period)
	res := wma(input, weights)
	alpha := 2.0/(float64(period) + 1)
	res[0] = input[0]
	// y_(i+1) = y_i + alpha * (x_(i+1) i y_i)
	for i:=0; i<len(input) - 1; i++ {
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


func wilder(input []float64, period int) []float64 {
	res := make([]float64, len(input))

	// Calculate the first WilderMA value
	sum := 0.0
	for _, value := range input[:period] {
		sum += value
	}
	res[period-1] = sum / float64(period)

	// Calculate the rest of the WilderMA values
	for i := period; i < len(input); i++ {
		res[i] = res[i-1] + (input[i] - res[i-1]) / float64(period)
	}

	return res
}