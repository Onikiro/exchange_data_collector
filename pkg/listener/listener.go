package listener

import (
	"exchange-data-collector/pkg/config"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Connection struct {
	Symbol   string
	Client   *influxdb2.Client
	WriteApi *api.WriteAPI
	StopC    chan struct{}
}

var connections = sync.Map{}

var Wg = sync.WaitGroup{}

func Listen(symbol string) {
	Wg.Add(1)
	config := config.LoadConfig()
	client, writeAPI := createInfluxClient(config)

	_, stopC, err := binance.WsBookTickerServe(symbol, handleWsEvent(writeAPI), handleErr())
	if err != nil {
		log.Fatalln(err)
	}

	connections.Store(symbol, Connection{
		Symbol:   symbol,
		Client:   client,
		WriteApi: writeAPI,
		StopC:    stopC,
	})
}

func StopListening(symbol string) {
	connection, ok := connections.Load(symbol)
	if !ok {
		log.Fatalln("Connection missed!")
	}

	v := connection.(Connection)
	v.StopC <- struct{}{}
	(*v.Client).Close()

	connections.Delete(symbol)
	Wg.Done()
}

func handleErr() func(err error) {
	return func(err error) {
		log.SetPrefix("ERROR: ")
		log.Println(err)
	}
}

func handleWsEvent(writeAPI *api.WriteAPI) func(event *binance.WsBookTickerEvent) {
	return func(event *binance.WsBookTickerEvent) {
		log.SetPrefix("INFO: ")
		log.Println(event)

		bestAskPrice, _ := strconv.ParseFloat(event.BestAskPrice, 64)
		bestBidPrice, _ := strconv.ParseFloat(event.BestBidPrice, 64)
		bestAskQty, _ := strconv.ParseFloat(event.BestAskQty, 64)
		bestBidQty, _ := strconv.ParseFloat(event.BestBidQty, 64)

		p := influxdb2.NewPoint("book_ticker",
			map[string]string{"symbol": event.Symbol},
			map[string]interface{}{"best_ask_price": bestAskPrice, "best_bid_price": bestBidPrice, "best_ask_qty": bestAskQty, "best_bid_qty": bestBidQty},
			time.Now().UTC())
		(*writeAPI).WritePoint(p)
		(*writeAPI).Flush()
	}
}

func createInfluxClient(config *config.Config) (*influxdb2.Client, *api.WriteAPI) {
	client := influxdb2.NewClient(config.InfluxUrl, config.InfluxToken)
	writeAPI := client.WriteAPI(config.InfluxOrg, config.InfluxBucket)
	return &client, &writeAPI
}
