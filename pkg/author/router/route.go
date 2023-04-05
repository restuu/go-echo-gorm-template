package router

import (
	"go-echo-gorm-tempate/pkg/author/service"

	"github.com/labstack/echo/v4"
)

// AuthorRouterService ...
type AuthorRouterService struct {
	UserGettingService service.AuthorGettingService
}

func NewAuthorRouter(e *echo.Echo, services AuthorRouterService) {
	g := e.Group("/authors")

	g.GET("", getAllAuthors(services.UserGettingService))
}

func getAllAuthors(svc service.AuthorGettingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		authors, err := svc.FindAll(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(200, authors)
	}
}
