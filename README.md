# Exchange Market Data Collector
CryptoExchanges market data collector, integrated with InfluxDB and SQLite

Features: 

* Writes market data from binance to InfluxDB

* Web API endpoints to add symbol to listen and save

* SQLite as database to save symbols

* Endpoint to read saved data from InfluxDB

Setup: 

1.  `docker volume create --name=dbdata`

2.  `docker volume create --name=influxdata`

3. `docker compose up -d`

4.  Setup influxDB in `localhost:8086`

5. Fill envs from influxDB in docker-compose.yaml

6. `docker-compose up -d --build --force-recreate`


TODO: 

- [x] Migrations (Create if not exists)

- [ ] Tests

- [x] Add docker support

- [x] Add docker-compose with sqlite db volume, influxdb and application
