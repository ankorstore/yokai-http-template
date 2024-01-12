package internal

import (
	"github.com/ankorstore/yokai-http-template/internal/service"
	"go.uber.org/fx"
)

func ProvideServices() fx.Option {
	return fx.Provide(
		service.NewWelcomeService,
	)
}
