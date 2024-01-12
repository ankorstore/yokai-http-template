package handler

import (
	"fmt"
	"net/http"

	"github.com/ankorstore/yokai-http-template/internal/service"
	"github.com/labstack/echo/v4"
)

type WelcomeHandler struct {
	service *service.WelcomeService
}

func NewWelcomeHandler(service *service.WelcomeService) *WelcomeHandler {
	return &WelcomeHandler{
		service: service,
	}
}

func (h *WelcomeHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Info("called WelcomeHandler")

		return c.String(http.StatusOK, fmt.Sprintf("Welcome to %s", h.service.Welcome(c.Request().Context())))
	}
}