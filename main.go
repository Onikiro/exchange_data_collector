package main

import (
	"exchange-data-collector/pkg/listener"
)

func main() {
	listener.Listen("BTCUSDT")
	listener.Listen("ETHUSDT")
	listener.Listen("BNBUSDT")

	listener.Wg.Wait()
}
