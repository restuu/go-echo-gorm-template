package main

import (
	"context"
	"errors"
	"fmt"
	"go-echo-gorm-tempate/adapter/config"
	"go-echo-gorm-tempate/app/author"
	"go-echo-gorm-tempate/app/book"
	"go-echo-gorm-tempate/domain"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Usecases struct {
	AuthorService domain.AuthorUsecase
	BookService   domain.BookUsecase
}

type Repositories struct {
	domain.AuthorRepository
}

type App struct {
	Context  context.Context
	Usecases Usecases
	Router   *echo.Echo
	Config   *config.Config
}

func (a *App) Start() {
	a.initRouter()
	a.setErrorHandler()
	a.setRoutes()
}
func (a *App) initRouter() {
	router := echo.New()
	a.Router = router

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.Recover())

}

func (a *App) setErrorHandler() {
	a.Router.HTTPErrorHandler = func(err error, c echo.Context) {
		var er *echo.HTTPError
		ok := errors.As(err, &er)
		if !ok {
			er = &echo.HTTPError{
				Code: http.StatusBadRequest,
				// might want to unwrap this error
				Message: err.Error(),
			}
		}

		if c.Response().Committed {
			return
		}

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(er.Code)
		} else {
			err = c.JSON(er.Code, er.Message)
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func (a *App) setRoutes() {
	book.SetBookRouter(a.Router, a.Usecases.BookService)
	author.SetAuthorRouter(a.Router, a.Usecases.AuthorService)
}

func (a *App) start() {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", a.Config.Port),
		BaseContext: func(_ net.Listener) context.Context {
			return a.Context
		},
	}

	// start server
	go func() {
		if err := a.Router.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.Router.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.Router.Shutdown(ctx); err != nil {
		a.Router.Logger.Fatal(err)
	}
}
