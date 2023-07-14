package brokers

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/d1l1x/gofin/utils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

var log = utils.NewZapLogger("broker", utils.Debug) //.Sugar()

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

	log.Info("Setup Broker", zap.String("name", "Alpaca"))

	if baseUrl == "" {
		log.Info("Use paper trading", zap.String("url", basePaperURL))
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
		log.Debug("Get credentials from environment")
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
		//timeToOpen := int(clock.NextOpen.Sub(clock.Timestamp).Minutes())
		//log.Debug("%d minutes until next market open\n", timeToOpen)
		log.Debug("Market is closed")
	}
	return false, nil
}

func (broker *AlpacaBroker) GetAccountInfo() {
	account, err := broker.trade.GetAccount()
	if err != nil {
		log.Fatal("Get account info", zap.Error(err))
	}
	log.Info("Account info", zap.Any("account", account))
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
		log.Info("Order placed", zap.String("type", "limit"), zap.String("id", order.ID))
	} else {
		log.Warn("Order failed", zap.String("type", "limit"), zap.Error(err))
	}
	return order.ID, nil
}

func (broker *AlpacaBroker) GetListOfAssets(status, class, exchange string) ([]alpaca.Asset, error) {
	if status == "" {
		status = "active"
	}
	if class == "" {
		class = "us_equity"
	}
	log.Debug("Get list of assets", zap.String("status", status), zap.String("class", class), zap.String("exchange", exchange))
	allAssets, err := broker.trade.GetAssets(alpaca.GetAssetsRequest{
		Status:     status,
		AssetClass: class,
		Exchange:   exchange,
	})
	if err != nil {
		return nil, fmt.Errorf("GetListOfAssets: %w", err)
	}
	assets := make([]alpaca.Asset, 0)
	log.Debug("Filter list of assets", zap.Strings("filters", []string{"tradable", "non-OTC"}))
	for _, asset := range allAssets {
		if !asset.Tradable || asset.Exchange == "OTC" {
			continue
		}
		assets = append(assets, asset)
	}
	return assets, nil
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
