package indicators

func RSL(in []float64, period int) ([]float64,error) {
	err := CheckInput(in, period)
	if err != nil {
		return nil, err
	}
	res := make([]float64, len(in))
	for i := period; i<len(in); i++ {
		res[i] = in[i]/Mean(in[i-period:i+1])
	}
	return res, nil
}