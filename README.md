# Exchange Market Data Collector
CryptoExchanges market data collector, integrated with InfluxDB and SQLite

Features: 

* Writes market data from binance to InfluxDB

* Web API endpoints to add symbol to listen and save

* SQLite as database to save symbols

* Endpoint to read saved data from InfluxDB

TODO: 

- [x] Migrations (Create if not exists)

- [ ] Tests

- [x] Add docker support

- [x] Add docker-compose with sqlite db volume, influxdb and application
