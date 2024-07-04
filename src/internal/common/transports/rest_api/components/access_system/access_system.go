package access_system

import (
	"context"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"time"
)

// accessSystem - компонент системы доступа http rest api.
type accessSystem struct {
	conf *Config
	ctx  context.Context

	components *components
	repository interface {
		GetUser(ctx context.Context, id types.ID) (us *common_entities.User, err error)

		GetRoute(ctx context.Context, method, path string) (route *common_entities.HttpRoute, err error)
		GetActiveRoute(ctx context.Context, method, path string) (route *common_entities.HttpRoute, err error)
		RegisterRoute(ctx context.Context, route *common_entities.HttpRoute) (err error)
		SetInactiveRoutes(ctx context.Context) (err error)

		GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error)
		RegisterToken(ctx context.Context, tok *common_entities.JwtToken) (err error)
	}
}

// components - компоненты компонента системы доступа http rest api.
type components struct {
	Logger logger.Logger
}

// RegisterRoutes - регистрация маршрутов в системе.
func (acc *accessSystem) RegisterRoutes(list ...*common_entities.HttpRouteConstructor) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

		trc.FunctionCall(list)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	for _, constructor := range list {
		var route = constructor.Build()

		route.RegisterTime = time.Now()

		if err = acc.repository.RegisterRoute(acc.ctx, route); err != nil {
			acc.components.Logger.Error().
				Format("Failed to register http rest api route: '%s'. ", err).Write()
			return
		}

		acc.components.Logger.Info().
			Format("The route '%s %s' is registered. ", route.Method, route.Path).
			Field("authorize", route.Authorize).Write()
	}

	return
}
