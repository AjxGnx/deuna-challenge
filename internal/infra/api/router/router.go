package router

import (
	_ "github.com/AjxGnx/deuna-challenge/docs"
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/handler"
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/router/group"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server           *echo.Echo
	transactionGroup group.Transaction
}

func New(
	server *echo.Echo,
	transactionGroup group.Transaction,
) *Router {
	return &Router{
		server,
		transactionGroup,
	}
}

func (router *Router) Init() {
	router.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	router.server.Use(middleware.Recover())

	basePath := router.server.Group("/api/exercise")

	basePath.GET("/swagger/*", echoSwagger.WrapHandler)
	basePath.GET("/health", handler.HealthCheck)

	basePath.Static("/static/", "./static")

	router.transactionGroup.Resource(basePath)
}
