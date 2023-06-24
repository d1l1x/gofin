package brokers

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/fatih/color"
	"github.com/shopspring/decimal"
	"os"
	"time"

	"github.com/d1l1x/gofin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const basePaperURL = "https://paper-api.alpaca.markets"

type AlpacaBroker struct {
	trade *alpaca.Client
	data  *marketdata.Client
}

type AlpacaCredentials struct {
	ApiKey    string
	ApiSecret string
}

func Alpaca(credentials *AlpacaCredentials, baseUrl string) *AlpacaBroker {

	log.SetFormatter(&gofin.PrefixedFormatter{
		Prefix:    "Alpaca",
		TextColor: color.FgGreen,
		Formatter: &logrus.TextFormatter{},
	})
	log.SetOutput(os.Stdout)

	if baseUrl == "" {
		baseUrl = basePaperURL
	}
	if credentials != nil {
		return &AlpacaBroker{
			// Alternatively you can set your key and secret using the
			// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
			trade: alpaca.NewClient(alpaca.ClientOpts{
				APIKey:    credentials.ApiKey,
				APISecret: credentials.ApiSecret,
				BaseURL:   baseUrl,
			}),
			data: marketdata.NewClient(marketdata.ClientOpts{
				APIKey:    credentials.ApiKey,
				APISecret: credentials.ApiSecret,
			}),
		}
	} else {
		log.Debugln("reading credentials from environment variables")
		return &AlpacaBroker{
			trade: alpaca.NewClient(alpaca.ClientOpts{
				BaseURL: baseUrl,
			}),
			data: marketdata.NewClient(marketdata.ClientOpts{}),
		}
	}
}

// IsMarketOpen checks if the market is currently open.
// If the market is open, it returns true and no error.
// If the market is closed, it logs the number of minutes until the market opens again,
// and returns false and no error.
// If there is an error while fetching the market clock information, it returns false and the error.
func (broker *AlpacaBroker) IsMarketOpen() (bool, error) {
	clock, err := broker.trade.GetClock()
	if err != nil {
		return false, fmt.Errorf("get clock: %w", err)
	}
	if clock.IsOpen {
		return true, nil
	} else {
		timeToOpen := int(clock.NextOpen.Sub(clock.Timestamp).Minutes())
		log.Infof("%d minutes until next market open\n", timeToOpen)
	}
	return false, nil
}

func (broker *AlpacaBroker) BuyingPower() (float64, error) {
	account, err := broker.trade.GetAccount()
	if err != nil {
		return 0, fmt.Errorf("get account: %w", err)
	}
	res, _ := account.BuyingPower.Float64()
	return res, nil
}

func (broker *AlpacaBroker) Cash() (float64, error) {
	account, err := broker.trade.GetAccount()
	if err != nil {
		return 0, fmt.Errorf("get account: %w", err)
	}
	res, _ := account.Cash.Float64()
	return res, nil
}

func (broker *AlpacaBroker) OpenPositions() ([]alpaca.Position, error) {
	positions, err := broker.trade.GetPositions()
	if err != nil {
		return nil, fmt.Errorf("list positions: %w", err)
	}
	return positions, nil
}

func (broker *AlpacaBroker) OpenOrders() ([]alpaca.Order, error) {
	orders, err := broker.trade.GetOrders(alpaca.GetOrdersRequest{
		Status: "open",
		Until:  time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("get open orders: %w", err)
	}
	return orders, nil
}

func (broker *AlpacaBroker) CancelOrder(orderId string) error {
	err := broker.trade.CancelOrder(orderId)
	if err != nil {
		return fmt.Errorf("cancel order: %w", err)
	}
	return nil
}

func (broker *AlpacaBroker) LimitOrder(side alpaca.Side, symbol string, quantity int, limitPrice float64, timeInForce alpaca.TimeInForce) (string, error) {
	qty := decimal.NewFromInt(int64(quantity))
	price := decimal.NewFromFloat(limitPrice)
	order, err := broker.trade.PlaceOrder(alpaca.PlaceOrderRequest{
		Symbol:      symbol,
		Qty:         &qty,
		Side:        side,
		Type:        alpaca.Limit,
		LimitPrice:  &price,
		TimeInForce: timeInForce,
	})
	if err == nil {
		log.Infof("limit order placed: %s", order.ID)
	} else {
		log.Warnf("limit order failed: %s", err.Error())
	}
	return order.ID, nil
}
