package access_system

import (
	"context"
	"github.com/gofiber/fiber/v3"
	http_routes_repository "sm-box/internal/services/authentication/components/access_system/repositories/http_routes"
	jwt_tokens_repository "sm-box/internal/services/authentication/components/access_system/repositories/jwt_tokens"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator = "system-[components]=access_system"
)

// AccessSystem - описание компонента системы доступа.
type AccessSystem interface {
	AuthenticationMiddlewareForRestAPI(ctx fiber.Ctx) (err error)
}

// New - создание компонента.
func New(ctx context.Context, conf *Config) (acc AccessSystem, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

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

	// Шлюзы
	{
		ref.gateways = new(gateways)
	}

	// Репозиторий
	{
		ref.repositories = new(repositories)

		if ref.repositories.HttpRoutes, err = http_routes_repository.New(ref.ctx, ref.conf.Repositories.HttpRoutes); err != nil {
			return
		}

		if ref.repositories.JwtTokens, err = jwt_tokens_repository.New(ref.ctx, ref.conf.Repositories.JwtTokens); err != nil {
			return
		}
	}

	acc = ref

	ref.components.Logger.Info().
		Text("The access system component has been created. ").
		Field("config", ref.conf).Write()

	return
}
