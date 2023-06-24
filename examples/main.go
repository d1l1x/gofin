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

	broker := brokers.Alpaca(nil, "")

	for {
		isOpen, err := broker.IsMarketOpen()
		if err != nil {
			log.Errorf("is market open: %v", err)
		}
		if !isOpen {
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

			positions, err := broker.OpenPositions()
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

			orders, err := broker.OpenOrders()
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

			orders, err = broker.OpenOrders()
			if err != nil {
				log.Errorf("number of open orders: %v", err)
			} else {
				log.Infof("number of open orders: %v", len(orders))
			}

			time.Sleep(1 * time.Minute)
			continue
		}
	}
}
