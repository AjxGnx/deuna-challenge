package providers

import (
	"github.com/AjxGnx/deuna-challenge/internal/app"
	"github.com/AjxGnx/deuna-challenge/internal/infra/adapters/db"
	"github.com/AjxGnx/deuna-challenge/internal/infra/adapters/db/repository"
	"github.com/AjxGnx/deuna-challenge/internal/infra/adapters/stripe"
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/handler"
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/router"
	"github.com/AjxGnx/deuna-challenge/internal/infra/api/router/group"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(db.ConnInstance)
	_ = Container.Provide(router.New)

	_ = Container.Provide(group.NewTransactionGroup)
	_ = Container.Provide(handler.NewTransactionHandler)
	_ = Container.Provide(app.NewTransactionApp)
	_ = Container.Provide(stripe.NewStripe)
	_ = Container.Provide(repository.NewTransactionRepo)

	return Container
}
