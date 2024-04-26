package system_access

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/app/transports/rest_api/components/system_access/repository"
	"sm-box/pkg/core/components/logger"
)

const (
	loggerInitiator = "transports-[http]-[rest_api]-[components]=system_access"
)

type SystemAccess interface {
	Middleware(ctx fiber.Ctx) (err error)
	RegisterRoutes(list ...*fiber.Route) (err error)

	BasicAuth(ctx fiber.Ctx) (err error)
}

func New(ctx context.Context, conf *Config) (sysAcc SystemAccess, err error) {
	var ref = &systemAccess{
		conf: conf,
		ctx:  ctx,
	}

	// Конфигурация
	{
		if err = ref.conf.FillEmptyFields().Validate(); err != nil {
			return
		}
	}

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозиторий
	{
		if ref.repository, err = repository.New(); err != nil {
			return
		}
	}

	sysAcc = ref

	return
}
