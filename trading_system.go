package gofin

import (
	"github.com/d1l1x/gofin/brokers"
	"github.com/d1l1x/gofin/providers"
	"github.com/d1l1x/gofin/utils"
	"go.uber.org/zap"
	"time"
	"context"
)

var log = utils.NewZapLogger("TradingSystem", utils.Debug) //.Sugar()

type TradingSystem struct {
	watchlist *utils.Watchlist
	broker    *brokers.AlpacaBroker
	cal       *utils.TradingCalendar
	provider  providers.DataProvider
	// mm 	 MoneyManager
	maxPositions int
}

func NewTradingSystem(broker *brokers.AlpacaBroker, watchlist *utils.Watchlist, cal *utils.TradingCalendar, provider providers.DataProvider) *TradingSystem {
	return &TradingSystem{
		watchlist: watchlist,
		broker:    broker,
		cal:       cal,
		provider:  provider,
	}
}

func (ts *TradingSystem) Init() {

	// TODO: Add error checking if necessary objects are initialized
	// i.e. broker, watchlist, calendar, etc.

	bp, err := ts.broker.BuyingPower()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("", zap.Float64("Buying power", bp))
}

func (ts *TradingSystem) CheckOpenPositions() int {

	log.Info("Get open positions")
	positions, err := ts.broker.GetOpenPositions()
	if err != nil {
		log.Error("Get open positions", zap.Error(err))
	}

	errPos := make(chan error, len(positions))
	for _, pos := range positions {
		go func(p brokers.Position) {
			log.Info("Check stopp loss", zap.String("Symbol", p.Symbol))
			// err := ts.broker.GetHistBars(a.Symbol, 100)
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			log.Info("Check take profit", zap.String("Symbol", p.Symbol))

			log.Info("Check time based exits", zap.String("Symbol", p.Symbol)

			errPos <- nil
		}(pos)
	}

	// Check possible channel errors
	for range ts.watchlist.Assets {
		if err := <-errPos; err != nil {
			log.Error("Channel error", zap.Error(err))
		}
	}

	return openPositions
}

func (ts *TradingSystem) WaitForTradingHours() {

	if !ts.cal.IsTradingDay(time.Now()) {
		log.Info("Not a trading day. Waiting for next trading day.")
		time.Sleep(time.Until(ts.cal.NextBusinessDay(time.Now())))
	}

	if time.Now().Before(ts.cal.OnClose.Start) {
		log.Info("Not on close. Waiting ...")
		time.Sleep(time.Until(ts.cal.OnClose.Start))
	}

}

func (ts *TradingSystem) processOrders(orders *[]brokers.Order, timeout time.Duration) (int,int) {

	log.Info("Process orders", zap.Int("Number of orders", len(*orders)), zap.Duration("Timeout", timeout))

	filled := make(chan *brokers.Order, len(*orders))
	allFilled := make(chan bool)
	errChan := make(chan error, len(*orders))

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // cancel when we are finished

	// submit orders
	go func() {
		for _, order := range *orders {
			go func(order brokers.Order) {
				qty,_ := order.Qty.Float64()
				limit, _ := order.LimitPrice.Float64()
				stop, _ := order.StopPrice.Float64()
				log.Info("Place order", zap.String("Symbol", order.Symbol), zap.Any("side", order.Side), zap.Float64("qty", qty), zap.Float64("limit", limit), zap.Float64("stop", stop))
				//TODO: Account for different order types
				orderID, err := sendOrderToBroker(order)
				if err != nil {
					errChan <- err
					return
				}
				// wait for order to be filled
				for {
					orderId, err := ts.broker.GetOrder(orderID)
					if err != nil {
						errChan <- err
						return
					}
					if orderId.Status == "filled" {
						break
					}
				}
				filled <- &order
			}(order)
		}
	}()

	// wait for all orders to be filled or for the deadline to be reached
	filledOrders := 0
	go func() {
		for i := 0; i < len(*orders); i++ {
			select {
			case <-filled:
				order := <-filled
				price, _ := order.FilledAvgPrice.Float64()
				qty, _ := order.FilledQty.Float64()
				log.Info("Order filled", zap.String("Symbol", order.Symbol), zap.Float64("qty", qty), zap.Float64("FillPrice", price))
				filledOrders++
			case err := <-errChan:
				log.Warn("Error sending order to broker", zap.Error(err))
			case <-ctx.Done():
				log.Info("Timeout reached. Cancel remaining orders")
				err := ts.broker.CancelAllOrders()
				if err != nil {
					log.Error(err.Error())
				}
				allFilled <- true
				return
			}
		}
		allFilled <- true
	}()

	<-allFilled

	cancelledOrders := len(*orders) - filledOrders

	return filledOrders, cancelledOrders
}

func (ts *TradingSystem) Run() {
	defer log.Sync() // flushes buffer, if any

	for {
		ts.WaitForTradingHours()

		openPositions := ts.CheckOpenPositions()

		if openPositions >= ts.maxPositions {
			log.Info("Max number of positions reached. Waiting for next trading day.")
			time.Sleep(time.Until(ts.cal.NextBusinessDay(time.Now())))
			continue
		}

		// Check assets for filter criteria
		errChan := make(chan error, len(ts.watchlist.Assets))

		var assetsToConsider []utils.Asset

		////TODO: Check setup

		// filter and rank all assets
		for i, asset := range ts.watchlist.Assets {
			go func(idx int, a utils.Asset) {
				bars, err := ts.provider.GetHistBars(a.Symbol, 100)
				if err != nil {
					errChan <- err
					return
				}
				passed := ts.watchlist.ApplyFilters(a.Symbol, bars)

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

		log.Info("Assets to consider", zap.Int("Number of assets", len(assetsToConsider)))

		log.Info("Prepare orders")

		log.Info("Check money management")

		var listOfOrders []brokers.Order

		timeout := ts.cal.OnClose.End.Sub(time.Now())
		filledOrders, cancelledOrders := ts.processOrders(&listOfOrders, timeout)

		//if all orders have been filled for today, wait for next trading day
		log.Info("Orders filled. Waiting for next trading day at close.", zap.Int("Filled orders", filledOrders), zap.Int("Cancelled orders", cancelledOrders))
		time.Sleep(time.Until(ts.cal.NextDayOnClose(time.Now()).Start))
	}
}
