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

// e.POST("/configs", postConfig)
func postConfig(c echo.Context) error {
	symbol := c.FormValue("symbol")

	db.AddConfig(symbol)
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
	db.DeleteConfig(symbol)
	listener.StopListening(symbol)
	return c.JSON(http.StatusOK, symbol)
}
