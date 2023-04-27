package router

import (
	"go-echo-gorm-tempate/pkg/book/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BookRouterService book router services dependencies.
type BookRouterService struct {
	BookGettingService service.BookGettingService
}

func NewBookRouter(e *echo.Echo, services BookRouterService) {
	g := e.Group("/books")

	g.GET("", getAllBooks(services.BookGettingService))
}

func getAllBooks(svc service.BookGettingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := svc.FindAll(c.Request().Context())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, books)
	}
}
