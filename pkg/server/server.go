package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"exchange-data-collector/pkg/db"
	"exchange-data-collector/pkg/listener"
)

func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/configs", postConfig)
	e.GET("/configs", getConfigs)
	e.DELETE("/configs/:symbol", deleteConfig)

	e.Logger.Fatal(e.Start(":1323"))
}

type Config struct {
	Symbol string
}

// e.POST("/configs", postConfig)
func postConfig(c echo.Context) error {
	config := new(Config)
	if err := c.Bind(config); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	symbol := config.Symbol
	if len(symbol) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	if !db.AddConfig(symbol) {
		return c.NoContent(http.StatusBadRequest)
	}

	listener.Listen(symbol)

	return c.JSON(http.StatusCreated, symbol)
}

// e.GET("/configs", getConfigs)
func getConfigs(c echo.Context) error {
	configs := db.GetConfigs()
	return c.JSON(http.StatusOK, configs)
}

// e.DELETE("/configs/:symbol", deleteConfig)
func deleteConfig(c echo.Context) error {
	symbol := c.Param("symbol")
	if !db.DeleteConfig(symbol) {
		return c.NoContent(http.StatusNotFound)
	}

	listener.StopListening(symbol)
	return c.JSON(http.StatusOK, symbol)
}
