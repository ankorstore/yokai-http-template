package internal

import (
	"go.uber.org/fx"
)

// ProvideServices is used to register the application services.
func ProvideServices() fx.Option {
	return fx.Options()
}
