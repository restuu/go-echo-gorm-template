package author

import (
	"go-echo-gorm-tempate/domain"

	"github.com/labstack/echo/v4"
)

func SetAuthorRouter(e *echo.Echo, uc domain.AuthorUsecase) {
	g := e.Group("/authors")

	g.GET("", getAllAuthors(uc))
}

func getAllAuthors(svc domain.AuthorUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		authors, err := svc.FindAll(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(200, authors)
	}
}
