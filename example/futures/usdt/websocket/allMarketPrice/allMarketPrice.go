package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usdt/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesAllMarketPriceWebsocketClient("!markPrice@arr")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		client.Subscribe(123, "!markPrice@arr@1s", "!markPrice@3s")
		client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.AllMarketPriceResponse:
			logger.Info("AllMarket Response: %v", response.(futuresusdt.AllMarketPriceResponse))
		case futuresusdt.AllMarketPriceCombinedResponse:
			logger.Info("AllMarketCombinedResponse: %v", response.(futuresusdt.AllMarketPriceCombinedResponse))
		case model.WebsocketCommonResponse:
			logger.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logger.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			logger.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)
	fmt.Scanln()

	client.Unsubscribe(123, "!markPrice@arr@1s", "!markPrice@3s")
	client.Close()
	logger.Info("Client closed")
}
