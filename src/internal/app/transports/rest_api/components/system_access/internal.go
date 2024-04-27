package system_access

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/common"
	"sm-box/internal/entities/transports/rest_api/route"
	"sm-box/internal/entities/transports/rest_api/token"
	"sm-box/pkg/core/components/logger"
)

type systemAccess struct {
	conf *Config
	ctx  context.Context

	components *components
	repository interface {
		UserInfo(ctx context.Context, id uint64) (us *common.User, err error)
		BasicAuth(ctx context.Context, username, password string) (us *common.User, tok *token.Token, err error)

		RouteInfo(ctx context.Context, method, path string) (info *route.Info, err error)
		RegisterRoute(ctx context.Context, info *route.Info) (err error)

		TokenInfo(ctx context.Context, data []byte) (tok *token.Token, err error)
		RegisterToken(ctx context.Context, tok *token.Token) (err error)
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
