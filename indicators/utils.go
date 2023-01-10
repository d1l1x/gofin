package indicators

func Mean(in []float64) float64 {
	res := 0.0
	for _, val := range in {
		res += val
	}
	return res/float64(len(in))
}
