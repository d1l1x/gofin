package gofin

import (
	"fmt"
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
		errChan := make(chan error, len(ts.watchlist.Assets))

		var assetsToConsider []utils.Asset

		for i, asset := range ts.watchlist.Assets {
			go func(idx int, a utils.Asset) {
				bars, err := ts.provider.GetHistBars(a.Symbol, 100)
				if err != nil {
					errChan <- err
					return
				}

				passed := ts.watchlist.ApplyFilters(a.Symbol, bars)

				////TODO: Check setup
				ts.watchlist.ApplyRanking(&ts.watchlist.Assets[idx], bars)
				if passed {
					assetsToConsider = append(assetsToConsider, ts.watchlist.Assets[idx])
				}
				errChan <- nil
			}(i, asset)
		}
		// Check possible channel errors
		for range ts.watchlist.Assets {
			if err := <-errChan; err != nil {
				log.Error("Channel error", zap.Error(err))
			}
		}

		log.Info("Rank assets")
		ts.watchlist.RankAssets(assetsToConsider)

		// print first 10 entries of assetsToConsider
		for i, asset := range assetsToConsider[:10] {
			fmt.Printf("First symbols: %v. %v, rank: %v\n", i, asset.Symbol, asset.Rank)
		}

		// print last 10 entries of assetsToConsider
		for i, asset := range assetsToConsider[len(assetsToConsider)-10:] {
			fmt.Printf("Last symbols: %v. %v, rank: %v\n", i, asset.Symbol, asset.Rank)
		}

		//for _, asset := range assetsToConsider.Assets {
		//	log.Info("Considering Symbol", zap.String("symbol", asset.Symbol))
		//}
		break
	}
}
