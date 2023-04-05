package main

import (
	userRouter "go-echo-gorm-tempate/pkg/author/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initializeRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	return e
}

func registerRouters(e *echo.Echo, services *services) {
	userRouter.NewAuthorRouter(e, services.AuthorRouterService)
}
