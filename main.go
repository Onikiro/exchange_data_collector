package main

import (
	"exchange-data-collector/pkg/db"
	"exchange-data-collector/pkg/listener"
	"exchange-data-collector/pkg/server"
)

func main() {
	db.ApplySchema()

	for _, v := range db.GetConfigs() {
		listener.Listen(v)
	}

	server.Start()
}
