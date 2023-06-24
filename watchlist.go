package watchlist

import (
	"errors"
	"github.com/d1l1x/gofin/indicators"
)

type Stock struct {
	symbol string
	bars   indicators.BarHistory
}

type Watchlist struct {
	Stocks  []*Stock
	Filters []*Filter
}

type Comparison int

const (
	LT = iota
	GT
	LE
	GE
)

type Filter struct {
	Indicator indicators.Indicator
	Operator  Comparison
	Value     interface{}
}

func (w *Watchlist) AddStock(s *Stock) {
	w.Stocks = append(w.Stocks, s)
}

func (w *Watchlist) SetStocks(stocks []*Stock) {
	w.Stocks = stocks
}

func (w *Watchlist) AddFilter(f *Filter) {
	w.Filters = append(w.Filters, f)
}

//func (w *Watchlist) ApplyFilters() *Watchlist {
//	filtered := &Watchlist{}
//	for _, s := range w.Stocks {
//		passes := true
//		for _, c := range w.Filters {
//			if !c(s) {
//				passes = false
//				break
//			}
//		}
//		if passes {
//			filtered.AddStock(s)
//		}
//	}
//	return filtered
//}

// NewFilter creates a new Filter with the given indicator, comparison operator, and value.
// The indicator should be an implementation of the Indicator interface, which has a Compute method
// that returns a slice of float64 values.
//
// The operator should be one of the Comparison constants: LT (less than), GT (greater than),
// LE (less than or equal to), or GE (greater than or equal to).
//
// The value can be either a float64 or an implementation of the Indicator interface. If it's a float64,
// it represents the value that the indicator's computed result will be compared to. If it's an Indicator,
// its Compute method will be called to get the value to compare to.
//
// The function returns a pointer to the newly created Filter.
//
// Usage:
//
//	filter := NewFilter(MyIndicator{}, LT, 100.0)
//	filter := NewFilter(MyIndicator{}, LT, OtherIndicator{})
//
// The first example creates a new Filter that uses MyIndicator as the indicator, LT as the comparison operator,
// and 100.0 as the value to compare to. The second example creates a new Filter that uses MyIndicator as the
// indicator, LT as the comparison operator, and the result of OtherIndicator's Compute method as the value to compare
// to.
func NewFilter(indicator indicators.Indicator, operator Comparison, value interface{}) *Filter {
	return &Filter{
		Indicator: indicator,
		Operator:  operator,
		Value:     value,
	}
}

// Apply computes the value of the Filter's Indicator and compares it to the Filter's value
// using the Filter's comparison operator. The method returns true if the comparison is true,
// and false otherwise.
//
// The Indicator's Compute method is called to get a slice of float64 values, and the last value
// in the slice is used for the comparison.
//
// If the Filter's value is a float64, it is used directly for the comparison. If the Filter's value
// is an Indicator, its Compute method is called to get a slice of float64 values, and the last value
// in the slice is used for the comparison.
//
// The comparison operator should be one of the Comparison constants: LT (less than), GT (greater than),
// LE (less than or equal to), or GE (greater than or equal to).
//
// If the Filter's value is neither a float64 nor an Indicator, or if the comparison operator is not
// a known Comparison, the method returns false.
//
// Usage:
//
//	result := filter.Apply()
//
// This will compute the Indicator's value, compare it to the Filter's value using the Filter's
// comparison operator, and return the result of the comparison.
func (f Filter) Apply() (bool, error) {
	val := f.Indicator.Compute()
	indicatorValue := val[len(val)-1]
	var compareValue float64

	// Check if the value to compare to is a float64 or an Indicator
	switch v := f.Value.(type) {
	case float64:
		compareValue = v
	case indicators.Indicator:
		res := v.Compute()
		compareValue = res[len(res)-1]
	default:
		return false, errors.New("the provided comparison value is neither a float64 nor an Indicator")
	}

	// Apply the comparison operation
	switch f.Operator {
	case LT:
		return indicatorValue < compareValue, nil
	case GT:
		return indicatorValue > compareValue, nil
	case LE:
		return indicatorValue <= compareValue, nil
	case GE:
		return indicatorValue >= compareValue, nil
	default:
		return false, errors.New("unknown comparison operator")
	}
}
