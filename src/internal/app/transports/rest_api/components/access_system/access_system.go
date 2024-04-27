package access_system

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/app/transports/rest_api/components/access_system/repository"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator = "transports-[http]-[rest_api]-[components]=access_system"
)

// AccessSystem - описание компонента системы доступа http rest api.
type AccessSystem interface {
	Middleware(ctx fiber.Ctx) (err error)
	RegisterRoutes(list ...*fiber.Route) (err error)

	BasicUserAuth(ctx fiber.Ctx) (err error)
}

// New - создание компонента системы доступа http rest api.
func New(ctx context.Context, conf *Config) (acc AccessSystem, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelComponent)

		trc.FunctionCall(ctx, conf)
		defer func() { trc.Error(err).FunctionCallFinished(acc) }()
	}

	var ref = &accessSystem{
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
		if ref.repository, err = repository.New(ctx, conf.Repository); err != nil {
			return
		}
	}

	acc = ref

	// Установить все маршруты в бд как не активные
	{
		if err = ref.repository.SetInactiveRoutes(ref.ctx); err != nil {
			return
		}
	}

	ref.components.Logger.Info().
		Text("The access system component has been created. ").
		Field("config", ref.conf).Write()

	return
}