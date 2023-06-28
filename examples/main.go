package main

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/d1l1x/gofin"
	"github.com/d1l1x/gofin/brokers"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log = logrus.New()

func main() {

	log.SetFormatter(&gofin.PrefixedFormatter{
		Prefix:    "Main",
		TextColor: color.FgRed,
		Formatter: &logrus.TextFormatter{},
	})
	log.SetOutput(os.Stdout)

	// Define trading windows
	calendar, err := gofin.NewTradingCalendarUS()
	if err != nil {
		log.Fatalf("trading calendar: %v", err)
	}
	nextDayOnOpen := calendar.NextDayOnOpen(time.Now())
	if err != nil {
		log.Fatalf("next day on open: %v", err)
	}
	nextDayOnClose := calendar.NextDayOnClose(time.Now())
	if err != nil {
		log.Fatalf("next day on close: %v", err)
	}

	log.Infof("next day on open: %v", nextDayOnOpen)
	log.Infof("next day on open: %v", nextDayOnClose)

	// Define broker
	broker := brokers.Alpaca(nil, "", logrus.DebugLevel)

	for {
		isOpen, err := broker.IsMarketOpen()
		if err != nil {
			log.Errorf("is market open: %v", err)
		}
		if isOpen {
			buyingPower, err := broker.BuyingPower()
			if err != nil {
				log.Errorf("buying power: %v", err)
			} else {
				log.Infof("buying power: %v", buyingPower)
			}

			cash, err := broker.Cash()
			if err != nil {
				log.Errorf("cash: %v", err)
			} else {
				log.Infof("cash: %v", cash)
			}

			positions, err := broker.GetOpenPositions()
			if err != nil {
				log.Errorf("number of open positions: %v", err)
			} else {
				log.Infof("number of open positions: %v", len(positions))
			}

			orderID, err := broker.LimitOrder(alpaca.Buy, "AAPL", 1, 100, alpaca.Day)
			if err != nil {
				log.Errorf("order couldn't been placed: %v", err)
			} else {
				log.Infof("order placed: %v", orderID)
			}

			time.Sleep(30 * time.Second)

			orders, err := broker.GetOpenOrders()
			if err != nil {
				log.Errorf("number of open orders: %v", err)
			} else {
				log.Infof("number of open orders: %v", len(orders))
			}

			for _, order := range orders {
				err := broker.CancelOrder(order.ID)
				if err != nil {
					log.Errorf("order couldn't been cancelled: %v", err)
				} else {
					log.Infof("order cancelled: %v %v", order.ID, order.Symbol)
				}
			}

			time.Sleep(30 * time.Second)

			orders, err = broker.GetOpenOrders()
			if err != nil {
				log.Errorf("number of open orders: %v", err)
			} else {
				log.Infof("number of open orders: %v", len(orders))
			}

			assets, err := broker.GetListOfAssets("", "", "", true)
			if err != nil {
				log.Errorf("list of assets: %v", err)
			} else {
				log.Infof("list of assets: %v", len(assets))
			}
			tradableAssets := 0
			for _, asset := range assets {
				if asset.Tradable {
					tradableAssets++
				}
			}
			log.Infof("tradable assets: %v", tradableAssets)

			shortableAssets := 0
			for _, asset := range assets {
				if asset.Shortable {
					shortableAssets++
				}
			}
			log.Infof("tradable assets: %v", shortableAssets)

			bars, err := broker.GetSymbolBars("AAPL", 200)
			if err != nil {
				log.Errorf("bars: %v", err)
			} else {
				for i, bar := range bars {
					log.Infof("bar %v: %v", i, bar)
				}
			}

			time.Sleep(1 * time.Minute)
			continue
		}
	}
}
