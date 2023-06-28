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

func Alpaca(credentials *AlpacaCredentials, baseUrl string, logLevel logrus.Level) *AlpacaBroker {

	log.SetFormatter(&gofin.PrefixedFormatter{
		Prefix:    "Alpaca",
		TextColor: color.FgGreen,
		Formatter: &logrus.TextFormatter{},
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)

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
		return false, fmt.Errorf("IsMarketOpen: %w", err)
	}
	if clock.IsOpen {
		return true, nil
	} else {
		timeToOpen := int(clock.NextOpen.Sub(clock.Timestamp).Minutes())
		log.Debugf("%d minutes until next market open\n", timeToOpen)
	}
	return false, nil
}

func (broker *AlpacaBroker) BuyingPower() (float64, error) {
	account, err := broker.trade.GetAccount()
	if err != nil {
		return 0, fmt.Errorf("BuyingPower: %w", err)
	}
	res, _ := account.BuyingPower.Float64()
	return res, nil
}

func (broker *AlpacaBroker) Cash() (float64, error) {
	account, err := broker.trade.GetAccount()
	if err != nil {
		return 0, fmt.Errorf("Cash: %w", err)
	}
	res, _ := account.Cash.Float64()
	return res, nil
}

func (broker *AlpacaBroker) GetOpenPositions() ([]alpaca.Position, error) {
	positions, err := broker.trade.GetPositions()
	if err != nil {
		return nil, fmt.Errorf("GetOpenPositions: %w", err)
	}
	return positions, nil
}

func (broker *AlpacaBroker) GetOpenPosition(symbol string) (*alpaca.Position, error) {
	position, err := broker.trade.GetPosition(symbol)
	if err != nil {
		return nil, fmt.Errorf("GetOpenPosition: %w", err)
	}
	return position, nil
}

func (broker *AlpacaBroker) GetOpenOrders() ([]alpaca.Order, error) {
	orders, err := broker.trade.GetOrders(alpaca.GetOrdersRequest{
		Status: "open",
		Until:  time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("GetOpenOrders: %w", err)
	}
	return orders, nil
}

func (broker *AlpacaBroker) CancelOrder(orderId string) error {
	err := broker.trade.CancelOrder(orderId)
	if err != nil {
		return fmt.Errorf("CancelOrder: %w", err)
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
		log.Infof("Limit order placed: %s", order.ID)
	} else {
		log.Warnf("Limit order failed: %s", err.Error())
	}
	return order.ID, nil
}

func (broker *AlpacaBroker) GetListOfAssets(status, class, exchange string, tradable bool) ([]alpaca.Asset, error) {
	if status == "" {
		status = "active"
	}
	if class == "" {
		class = "us_equity"
	}
	allAssets, err := broker.trade.GetAssets(alpaca.GetAssetsRequest{
		Status:     status,
		AssetClass: class,
		Exchange:   exchange,
	})
	tradableAsset := make([]alpaca.Asset, 0)
	if tradable {
		for _, asset := range allAssets {
			if !asset.Tradable {
				continue
			}
			tradableAsset = append(tradableAsset, asset)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("GetListOfAssets: %w", err)
	}
	return tradableAsset, nil
}

func (broker *AlpacaBroker) GetSymbolBars(symbol string, period int) ([]marketdata.Bar, error) {
	//businessDay := broker.calendar.calendar.lastBusinessDay(time.Now())
	bars, err := broker.data.GetBars(symbol, marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneDay,
		Start:     time.Now().AddDate(0, 0, -period),
		End:       time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("GetSymbolBars: %s", err.Error())
	}
	return bars, nil
}
