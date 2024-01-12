package service_test

import (
	"context"
	"testing"

	"github.com/ankorstore/yokai-http-template/internal"
	"github.com/ankorstore/yokai-http-template/internal/service"
	"github.com/ankorstore/yokai/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestWelcome(t *testing.T) {
	var cfg *config.Config

	internal.RunTest(t, fx.Populate(&cfg))

	ctx := context.Background()
	svc := service.NewWelcomeService(cfg)

	assert.Equal(t, "app name: http-app, app version: 0.1.0", svc.Welcome(ctx))
}
