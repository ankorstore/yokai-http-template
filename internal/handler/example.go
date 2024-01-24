package handler

import (
	"fmt"
	"net/http"

	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
)

// ExampleHandler is an example HTTP handler.
type ExampleHandler struct {
	config *config.Config
}

// NewExampleHandler returns a new [ExampleHandler].
func NewExampleHandler(config *config.Config) *ExampleHandler {
	return &ExampleHandler{
		config: config,
	}
}

// Handle handles HTTP requests.
func (h *ExampleHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Welcome to %s.", h.config.AppName()))
	}
}
