package indicators

func Last(bars BarHistory, valueType LastValueType) *LastValue {
	return &LastValue{
		BarHistoryIndicator: NewBarHistoryIndicator(bars, 0),
		ValueType:           valueType,
	}
}

type LastValueType int

const (
	Open = iota
	High
	Low
	Close
	Volume
)

type LastValue struct {
	BarHistoryIndicator
	ValueType LastValueType
}

func (ind *LastValue) Compute() float64 {
	switch ind.ValueType {
	case Open:
		return ind.Input.Open[len(ind.Input.Open)-1]
	case High:
		return ind.Input.High[len(ind.Input.High)-1]
	case Low:
		return ind.Input.Low[len(ind.Input.Low)-1]
	case Close:
		return ind.Input.Close[len(ind.Input.Close)-1]
	case Volume:
		return float64(ind.Input.Volume[len(ind.Input.Volume)-1])
	}
	return 0
}
