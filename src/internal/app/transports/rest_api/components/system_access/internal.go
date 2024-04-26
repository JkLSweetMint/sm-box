package system_access

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/pkg/core/components/logger"
)

type systemAccess struct {
	conf *Config
	ctx  context.Context

	components *components
	repository interface {
	}
}

type components struct {
	Logger logger.Logger
}

func (s *systemAccess) Middleware(ctx fiber.Ctx) (err error) {
	return
}

func (s *systemAccess) RegisterRoutes(list ...*fiber.Route) (err error) {
	return
}

func (s *systemAccess) BasicAuth(ctx fiber.Ctx) (err error) {
	return
}
