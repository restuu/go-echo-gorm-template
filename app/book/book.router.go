package book

import (
	"go-echo-gorm-tempate/domain"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	service domain.BookUsecase
}

func SetBookRouter(e *echo.Echo, services domain.BookUsecase) {
	g := e.Group("/books")

	h := &bookHandler{service: services}

	g.GET("", h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, books)
}
