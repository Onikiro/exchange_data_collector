package reader

import (
	"context"
	"errors"
	"exchange-data-collector/pkg/config"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type MeasurementData struct {
	From   string
	To     string
	Symbol string
	Data   []DataPoint
}

type DataPoint struct {
	BestAskPrice float64
	BestBidPrice float64
	BestAskQty   float64
	BestBidQty   float64
	Time         string
}

func Read(from string, to string, symbol string) (*MeasurementData, error) {

	_, err := time.Parse(time.RFC3339, from)
	if err != nil {
		log.SetPrefix("ERROR: ")
		log.Println(err)
		return nil, errors.New("time format is not rfc3339")
	}

	_, err = time.Parse(time.RFC3339, to)
	if err != nil {
		log.SetPrefix("ERROR: ")
		log.Println(err)
		return nil, errors.New("time format is not rfc3339")
	}

	config := config.LoadConfig()
	client := influxdb2.NewClient(config.InfluxUrl, config.InfluxToken)
	defer client.Close()
	queryAPI := client.QueryAPI(config.InfluxOrg)

	query := fmt.Sprintf("from(bucket:\"%s\")|> range(start: %s, stop: %s) |> filter(fn: (r) => r._measurement == \"book_ticker\" and r[\"symbol\"] == \"%s\") |> group()", config.InfluxBucket, from, to, symbol)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.SetPrefix("ERROR: ")
		log.Println(err)
	}

	data := make(map[string]map[string]float64, 32)
	for result.Next() {

		time := result.Record().Time().String()
		if _, ok := data[time]; !ok {
			data[time] = make(map[string]float64)
		}

		switch result.Record().ValueByKey("_field") {
		case "best_ask_price":
			data[time]["best_ask_price"] = result.Record().Value().(float64)
		case "best_bid_price":
			data[time]["best_bid_price"] = result.Record().Value().(float64)
		case "best_ask_qty":
			data[time]["best_ask_qty"] = result.Record().Value().(float64)
		case "best_bid_qty":
			data[time]["best_bid_qty"] = result.Record().Value().(float64)
		}
	}

	points := make([]DataPoint, len(data))
	i := 0
	for time, v := range data {
		points[i] = DataPoint{
			Time:         time,
			BestAskPrice: v["best_ask_price"],
			BestBidPrice: v["best_bid_price"],
			BestAskQty:   v["best_ask_qty"],
			BestBidQty:   v["best_bid_qty"],
		}
		i++
	}

	if result.Err() != nil {
		log.SetPrefix("ERROR: ")
		log.Println(err)
	}

	return &MeasurementData{
		From:   from,
		To:     to,
		Symbol: symbol,
		Data:   points,
	}, nil
}
