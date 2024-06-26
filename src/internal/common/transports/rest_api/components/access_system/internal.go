package access_system

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"strings"
	"time"
)

// accessSystem - компонент системы доступа http rest api.
type accessSystem struct {
	conf *Config
	ctx  context.Context

	components *components
	repository interface {
		GetUser(ctx context.Context, id types.ID) (us *entities.User, err error)

		GetRoute(ctx context.Context, method, path string) (route *entities.HttpRoute, err error)
		GetActiveRoute(ctx context.Context, method, path string) (route *entities.HttpRoute, err error)
		RegisterRoute(ctx context.Context, route *entities.HttpRoute) (err error)
		SetInactiveRoutes(ctx context.Context) (err error)

		GetToken(ctx context.Context, data string) (tok *entities.JwtToken, err error)
		RegisterToken(ctx context.Context, tok *entities.JwtToken) (err error)
	}
}

// components - компоненты компонента системы доступа http rest api.
type components struct {
	Logger logger.Logger
}

// RegisterRoutes - регистрация маршрутов в системе.
func (acc *accessSystem) RegisterRoutes(list ...*fiber.Route) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

		trc.FunctionCall(list)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	for _, r := range list {
		var route = new(entities.HttpRoute).FillEmptyFields()

		route.Active = true
		route.Method = strings.ToUpper(r.Method)
		route.Path = r.Path
		route.RegisterTime = time.Now()

		if err = acc.repository.RegisterRoute(acc.ctx, route); err != nil {
			acc.components.Logger.Error().
				Format("Failed to register http rest api route: '%s'. ", err).Write()
			return
		}
	}

	return
}
