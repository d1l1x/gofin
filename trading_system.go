package gofin

import (
	"github.com/d1l1x/gofin/brokers"
	"github.com/d1l1x/gofin/providers"
	"github.com/d1l1x/gofin/utils"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"time"
)

var log = utils.NewZapLogger("TradingSystem", utils.Debug) //.Sugar()

type TradingSystem struct {
	watchlist *utils.Watchlist
	broker    *brokers.AlpacaBroker
	cal       *utils.TradingCalendar
	provider  providers.DataProvider
	// mm 	 MoneyManager
}

func NewTradingSystem(broker *brokers.AlpacaBroker, watchlist *utils.Watchlist, cal *utils.TradingCalendar, provider providers.DataProvider) *TradingSystem {
	return &TradingSystem{
		watchlist: watchlist,
		broker:    broker,
		cal:       cal,
		provider:  provider,
	}
}

func (ts *TradingSystem) Init(logLevel logrus.Level) {

	// TODO: Add error checking if necessary objects are initialized
	// i.e. broker, watchlist, calendar, etc.

	bp, err := ts.broker.BuyingPower()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("", zap.Float64("Buying power", bp))
}

func (ts *TradingSystem) Run() {

	defer log.Sync() // flushes buffer, if any

	for {

		if !ts.cal.IsTradingDay(time.Now()) {
			log.Info("Not a trading day. Waiting for next trading day.")
			time.Sleep(1 * time.Hour)
			continue
		}

		//TODO: Check open positions
		//TODO: Check stopp Loss
		//TODO: Take Profit
		//TODO: Check time based exits

		// Check assets for filter criteria
		passed := make(chan bool, len(ts.watchlist.Assets))
		errChan := make(chan error, len(ts.watchlist.Assets))

		for _, asset := range ts.watchlist.Assets {
			go func(symbol string) {
				bars, err := ts.provider.GetHistBars(symbol, 100)
				if err != nil {
					errChan <- err
					return
				}

				log.Debug("Apply filters", zap.String("symbol", symbol))
				passed <- ts.watchlist.ApplyFilters(bars)
				//TODO: Add computation of rank

				errChan <- nil
			}(asset.Symbol)
		}

		assetsToConsider := utils.Watchlist{}

		for _, asset := range ts.watchlist.Assets {
			select {
			case pass := <-passed:
				if pass {
					assetsToConsider.AddAsset(asset)
				}
			case err := <-errChan:
				if err != nil {
					log.Error("Channel error", zap.Error(err))
				}
			}
		}

		for _, asset := range assetsToConsider.Assets {
			log.Info("Considering Symbol", zap.String("symbol", asset.Symbol))
		}
		break
	}
}
