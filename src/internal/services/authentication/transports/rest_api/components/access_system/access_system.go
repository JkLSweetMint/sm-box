package access_system

import (
	"context"
	"github.com/gofiber/fiber/v3"
	src_access_system "sm-box/internal/common/transports/rest_api/components/access_system"
	"sm-box/internal/common/transports/rest_api/components/access_system/repository"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator = "transports-[http]-[rest_api]-[components]=access_system"
)

// AccessSystem - описание компонента системы доступа http rest api.
type AccessSystem interface {
	src_access_system.AccessSystem

	BasicUserAuth(ctx fiber.Ctx) (err error)
}

// New - создание компонента системы доступа http rest api.
func New(ctx context.Context, conf *src_access_system.Config) (acc AccessSystem, err error) {
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

	// Исходный
	{
		if ref.AccessSystem, err = src_access_system.New(ctx, conf); err != nil {
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

	ref.components.Logger.Info().
		Text("The access system component has been created. ").
		Field("config", ref.conf).Write()

	return
}
