package utils

import (
	"github.com/d1l1x/gofin/indicators"
	"testing"
)

type MockIndicator struct {
	Input  []float64
	Values []float64
	Bars   *indicators.BarHistory
}

func NewMockIndicator(values []float64) *MockIndicator {
	bars := &indicators.BarHistory{
		Open:   []float64{},
		High:   []float64{},
		Low:    []float64{},
		Close:  values,
		Volume: []int64{},
	}
	return &MockIndicator{Values: values, Bars: bars}
}

func (m *MockIndicator) Compute() []float64 {
	return m.Values
}

func (m *MockIndicator) SetInput(bars *indicators.BarHistory) {
	m.Input = bars.Close
}

func TestWatchAddAsset(t *testing.T) {
	w := &Watchlist{}
	w.AddAsset(Asset{Symbol: "AAPL"})
	w.AddAsset(Asset{Symbol: "MSFT"})

	if len(w.Assets) != 2 {
		t.Errorf("Expected 2 stocks, got %d", len(w.Assets))
	}

	if w.Assets[0].Symbol != "AAPL" {
		t.Errorf("Expected AAPL, got %s", w.Assets[0].Symbol)
	}

	if w.Assets[1].Symbol != "MSFT" {
		t.Errorf("Expected MSFT, got %s", w.Assets[1].Symbol)
	}

}

func TestAddFilterFloat(t *testing.T) {
	w := &Watchlist{}
	filter := NewFilter(&MockIndicator{}, LT, 0.0)
	w.AddFilter(filter)

	if len(w.Filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(w.Filters))
	}
}

func TestAddFilterIndicator(t *testing.T) {
	w := &Watchlist{}
	filter := NewFilter(&MockIndicator{}, LT, MockIndicator{})
	w.AddFilter(filter)

	if len(w.Filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(w.Filters))
	}
}

func TestAddFilterFloatAndIndicator(t *testing.T) {
	w := &Watchlist{}
	filter := NewFilter(&MockIndicator{}, LT, 0.0)
	w.AddFilter(filter)
	filter = NewFilter(&MockIndicator{}, LT, MockIndicator{})
	w.AddFilter(filter)

	if len(w.Filters) != 2 {
		t.Errorf("Expected 2 filters, got %d", len(w.Filters))
	}
}

func TestFilterFloatApplyLT(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, LT, 3.0)

	res, _ := filter.apply(mock.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, LT, 3.0)

	res, _ = filter.apply(mock.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyLT(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter := NewFilter(mock1, LT, mock2)

	res, _ := filter.apply(mock1.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 5})

	filter = NewFilter(mock1, LT, mock2)

	res, _ = filter.apply(mock1.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterFloatApplyGT(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, GT, 6.0)

	res, _ := filter.apply(mock.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, GT, 1.0)

	res, _ = filter.apply(mock.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyGT(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 6})

	filter := NewFilter(mock1, GT, mock2)

	res, _ := filter.apply(mock1.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 1})

	filter = NewFilter(mock1, GT, mock2)

	res, _ = filter.apply(mock1.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterFloatApplyLE(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, LE, 4.0)

	res, _ := filter.apply(mock.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, LE, 2.0)

	res, _ = filter.apply(mock.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyLE(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 4})

	filter := NewFilter(mock1, LE, mock2)

	res, _ := filter.apply(mock1.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter = NewFilter(mock1, LE, mock2)

	res, _ = filter.apply(mock1.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}
func TestFilterFloatApplyGE(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, GE, 6.0)

	res, _ := filter.apply(mock.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, GE, 2.0)

	res, _ = filter.apply(mock.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyGE(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 6})

	filter := NewFilter(mock1, GE, mock2)

	res, _ := filter.apply(mock1.Bars)
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter = NewFilter(mock1, GE, mock2)

	res, _ = filter.apply(mock1.Bars)
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterApplyUnknownComp(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, 4, 6.0)

	_, err := filter.apply(mock.Bars)
	if err == nil {
		t.Errorf("Expected filter to return error, got nil")
	}
}
