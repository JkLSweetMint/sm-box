package http_access_system

import (
	"context"
	"github.com/gofiber/fiber/v3"
	jwt_tokens_repository "sm-box/internal/services/authentication/components/http_access_system/repositories/jwt_tokens"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
	users_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/users_service"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator = "system-[components]=access_system"
)

// AccessSystem - описание компонента системы доступа.
type AccessSystem interface {
	BasicAuthentication(ctx fiber.Ctx) (err error)
}

// New - создание компонента.
func New(ctx context.Context, conf *Config) (acc AccessSystem, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(acc) }()
	}

	var ref = &accessSystem{
		ctx:  ctx,
		conf: conf,
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

		// Projects
		{
			if ref.gateways.Projects, err = projects_service_gateway.New(ctx); err != nil {
				return
			}
		}

		// Users
		{
			if ref.gateways.Users, err = users_service_gateway.New(ctx); err != nil {
				return
			}
		}
	}

	// Репозиторий
	{
		ref.repositories = new(repositories)

		// JwtTokens
		{
			if ref.repositories.JwtTokens, err = jwt_tokens_repository.New(ctx, ref.conf.Repositories.JwtTokens); err != nil {
				return
			}
		}
	}

	acc = ref

	ref.components.Logger.Info().
		Text("The access system component has been created. ").
		Field("config", ref.conf).Write()

	return
}
