package service

import (
	"context"
	"fmt"

	"github.com/ankorstore/yokai/config"
)

type WelcomeService struct {
	config *config.Config
}

func NewWelcomeService(config *config.Config) *WelcomeService {
	return &WelcomeService{
		config: config,
	}
}

func (s *WelcomeService) Welcome(ctx context.Context) string {
	return fmt.Sprintf("app name: %s, app version: %s", s.config.AppName(), s.config.AppVersion())
}
