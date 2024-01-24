package internal

import (
	"github.com/ankorstore/yokai-worker-template/internal/handler"
	"github.com/ankorstore/yokai/fxhttpserver"
	"go.uber.org/fx"
)

// ProvideRouting is used to register the application HTTP handlers.
func ProvideRouting() fx.Option {
	return fx.Options(
		fxhttpserver.AsHandler("GET", "", handler.NewExampleHandler),
	)
}
