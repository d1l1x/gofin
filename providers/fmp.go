package providers

import (
	"context"
	"fmt"
	"github.com/d1l1x/gofin/indicators"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"os"
	"time"
)

type FmpProvider struct {
	Client  *fmp.APIClient
	Limiter *rate.Limiter
}

func FMP(apiKey string, requestLimit int) *FmpProvider {

	if apiKey == "" {
		log.Debug("Reading credentials from environment variables")
		apiKey = os.Getenv("FMP_API_KEY")
	}

	// Init your custome API client
	log.Debug("Initializing FMP API client")
	client, err := fmp.NewAPIClient(fmp.Config{
		APIKey:  apiKey, // Set Your API Key from site, default: demo
		Debug:   false,  // Set flag for debug request and response, default: false
		Timeout: 60,     // Set timeout for http client, default: 25
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("Initializing rate limiter")
	limiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(requestLimit)), 1)

	return &FmpProvider{Client: client, Limiter: limiter}
}

func (fmp *FmpProvider) GetHistBars(symbol string, period int) (*indicators.BarHistory, error) {

	// Wait for a token from the rate limiter
	if err := fmp.Limiter.Wait(context.Background()); err != nil {
		return nil, fmt.Errorf("Not allowed to proceed for symbol %s: %v\n", symbol, err)
	}

	log.Debug("Get historical bars", zap.String("symbol", symbol))
	bars, err := fmp.Client.Stock.DailyLastNDays(symbol, period)
	if err != nil {
		return nil, err
	}
	if len(bars.Historical) >= period {
		//log.Debug("Got bars: %v [%v, ...]", symbol, bars.Historical[0].AdjClose)
		history := new(indicators.BarHistory)
		for _, bar := range bars.Historical {
			history.Open = append(history.Open, bar.Open)
			history.High = append(history.High, bar.High)
			history.Low = append(history.Low, bar.Low)
			//TODO: Decide whether to use adjusted close or close
			history.Close = append(history.Close, bar.Close)
			history.Volume = append(history.Volume, int64(bar.Volume))
		}
		return history, nil
	} else {
		return nil, fmt.Errorf("Not enough bars for %v", symbol)
	}
}
