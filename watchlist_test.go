package watchlist

import (
	"github.com/d1l1x/gofin/indicators"
	"testing"
)

type MockIndicator struct {
	indicators.Indicator
	Values []float64
}

func NewMockIndicator(values []float64) *MockIndicator {
	return &MockIndicator{Values: values}
}

func (m *MockIndicator) Compute() []float64 {
	return m.Values
}

func TestWatchAddStock(t *testing.T) {
	w := &Watchlist{}
	w.AddStock(&Stock{symbol: "AAPL", bars: indicators.BarHistory{}})
	w.AddStock(&Stock{symbol: "MSFT", bars: indicators.BarHistory{}})

	if len(w.Stocks) != 2 {
		t.Errorf("Expected 2 stocks, got %d", len(w.Stocks))
	}

	if w.Stocks[0].symbol != "AAPL" {
		t.Errorf("Expected AAPL, got %s", w.Stocks[0].symbol)
	}

	if w.Stocks[1].symbol != "MSFT" {
		t.Errorf("Expected MSFT, got %s", w.Stocks[1].symbol)
	}

}

func TestSetStocks(t *testing.T) {
	w := &Watchlist{}
	w.SetStocks([]*Stock{
		{symbol: "AAPL", bars: indicators.BarHistory{}},
		{symbol: "MSFT", bars: indicators.BarHistory{}},
	})

	if len(w.Stocks) != 2 {
		t.Errorf("Expected 2 stocks, got %d", len(w.Stocks))
	}

	if w.Stocks[0].symbol != "AAPL" {
		t.Errorf("Expected AAPL, got %s", w.Stocks[0].symbol)
	}

	if w.Stocks[1].symbol != "MSFT" {
		t.Errorf("Expected MSFT, got %s", w.Stocks[1].symbol)
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

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, LT, 3.0)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyLT(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter := NewFilter(mock1, LT, mock2)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 5})

	filter = NewFilter(mock1, LT, mock2)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterFloatApplyGT(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, GT, 6.0)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, GT, 1.0)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyGT(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 6})

	filter := NewFilter(mock1, GT, mock2)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 1})

	filter = NewFilter(mock1, GT, mock2)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterFloatApplyLE(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, LE, 4.0)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, LE, 2.0)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyLE(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 4})

	filter := NewFilter(mock1, LE, mock2)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter = NewFilter(mock1, LE, mock2)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}
func TestFilterFloatApplyGE(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, GE, 6.0)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	filter = NewFilter(mock, GE, 2.0)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterIndicatorApplyGE(t *testing.T) {
	mock1 := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	mock2 := NewMockIndicator([]float64{1, 2, 3, 4, 6})

	filter := NewFilter(mock1, GE, mock2)

	res, _ := filter.Apply()
	if res {
		t.Errorf("Expected filter to return false, got true")
	}

	mock1 = NewMockIndicator([]float64{1, 2, 3, 4, 2})
	mock2 = NewMockIndicator([]float64{1, 2, 3, 4, 2})

	filter = NewFilter(mock1, GE, mock2)

	res, _ = filter.Apply()
	if !res {
		t.Errorf("Expected filter to return true, got false")
	}

}

func TestFilterApplyUnknownComp(t *testing.T) {
	mock := NewMockIndicator([]float64{1, 2, 3, 4, 5})
	filter := NewFilter(mock, 4, 6.0)

	_, err := filter.Apply()
	if err == nil {
		t.Errorf("Expected filter to return error, got nil")
	}
}
