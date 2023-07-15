package utils

import (
	"errors"
	"github.com/d1l1x/gofin/indicators"
	"go.uber.org/zap"
	"sort"
)

type Asset struct {
	Symbol string
	Name   string
	Id     string
	Rank   float64
}

type Watchlist struct {
	Assets  []Asset
	Filters []Filter
	Ranking *Ranking
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

func NewWatchlist(assets []Asset, filters []Filter, ranking *Ranking) *Watchlist {
	log.Debug("Create new watchlist")
	return &Watchlist{
		Assets:  assets,
		Filters: filters,
		Ranking: ranking,
	}
}

func (w *Watchlist) AddAsset(a Asset) {
	w.Assets = append(w.Assets, a)
}

func (w *Watchlist) AddFilter(f Filter) {
	log.Debug("Add filter to watchlist")
	w.Filters = append(w.Filters, f)
}

func (w *Watchlist) ApplyFilters(symbol string, bars *indicators.BarHistory) bool {
	log.Debug("Apply filters", zap.String("symbol", symbol))
	for _, filter := range w.Filters {
		res, _ := filter.apply(bars)
		if !res {
			return false
		}
	}
	return true
}

type RankOrder int

const (
	Ascending = iota
	Descending
)

type Ranking struct {
	Indicator indicators.Indicator
	Order     RankOrder
}

func (w *Watchlist) AddRanking(r *Ranking) {
	log.Debug("Add ranking to watchlist")
	w.Ranking = r
}

func (w *Watchlist) ApplyRanking(asset *Asset, bars *indicators.BarHistory) {
	if w.Ranking != nil {
		w.Ranking.Indicator.SetInput(bars)
		val := w.Ranking.Indicator.Compute()
		asset.Rank = val[len(val)-1]
	}
}

func (w *Watchlist) RankAssets(assets []Asset) {
	if w.Ranking != nil {
		switch w.Ranking.Order {
		case Ascending:
			sort.Slice(assets, func(i, j int) bool { return (assets)[i].Rank < (assets)[j].Rank })
		case Descending:
			sort.Slice(assets, func(i, j int) bool { return (assets)[i].Rank > (assets)[j].Rank })
		}
	}
}

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
	log.Debug("New Filter", zap.Any("indicator", indicator), zap.Any("operator", operator), zap.Any("value", value))
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
// TODO: Put error checking into NewFilter function
func (f Filter) apply(bars *indicators.BarHistory) (bool, error) {

	f.Indicator.SetInput(bars)

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
