package providers

import (
	"github.com/d1l1x/gofin/indicators"
	"github.com/d1l1x/gofin/utils"
)

type DataProvider interface {
	GetHistBars(symbol string, period int) (*indicators.BarHistory, error)
}

var log = utils.NewZapLogger("data_provider", utils.Debug)
