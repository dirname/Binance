package main

import (
	"fmt"
	"github.com/dirname/binance/futures/usdt/websocket/market"
	logger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	"time"
)

func main() {
	client := futuresusdt.NewUSDTFuturesLiquidationOrderWebsocketClient("xrpusdt@forceOrder")
	client.SetReadTimerInterval(5 * time.Second)
	client.SetReconnectWaitTime(5 * time.Second)
	client.SetHandler(func() {
		//client.Subscribe(123, "xrpusdt@forceOrder", "egldusdt@forceOrder")
		//client.SetCombined(true, 123)
	}, func(response interface{}) {
		switch response.(type) {
		case futuresusdt.LiquidationOrderResponse:
			logger.Info("LiquidationOrder Response: %v", response.(futuresusdt.LiquidationOrderResponse))
		case futuresusdt.LiquidationOrderCombinedResponse:
			logger.Info("LiquidationOrderCombinedResponse: %v", response.(futuresusdt.LiquidationOrderCombinedResponse))
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

	//client.Unsubscribe(123, "xrpusdt@forceOrder", "egldusdt@forceOrder")
	//client.Close()
	logger.Info("Client closed")
}
