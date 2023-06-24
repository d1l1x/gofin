package brokers

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"

	log "github.com/sirupsen/logrus"
)

const basePaperURL = "https://paper-api.alpaca.markets"

type AlpacaBroker struct {
	trade *alpaca.Client
	data  *marketdata.Client
}

func (broker *AlpacaBroker) Alpaca(apiKey, apiSecret, baseUrl string) *AlpacaBroker {
	if baseUrl == "" {
		baseUrl = basePaperURL
	}
	return &AlpacaBroker{
		trade: alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
			BaseURL:   baseUrl,
		}),
		data: marketdata.NewClient(marketdata.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
		}),
	}
}

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
