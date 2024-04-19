package main

import (
	"fmt"

	"github.com/AjxGnx/deuna-challenge/internal/infra/api/router"

	"github.com/AjxGnx/deuna-challenge/cmd/providers"
	"github.com/AjxGnx/deuna-challenge/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// @title        Deuna Challenge
// @version      1.0.0
// @description  Deuna Challenge Manager
// @license.name Alirio Gutierrez
// @BasePath     /api/exercise
// @schemes      http
func main() {
	container := providers.BuildContainer()
	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%v", config.Environments().ServerHost,
			config.Environments().ServerPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
