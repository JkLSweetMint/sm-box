package http_rest_api

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	cors_middleware "github.com/gofiber/fiber/v3/middleware/cors"
	"path"
	common_errors "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/rest_api/io"
)

// initFiberServer - инициализация http сервера fiber.
func (srv *server) initFiberServer() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Starting the initialization of the fiber http server... ").
		Field("config", srv.conf.Server.ToFiberConfig()).Write()

	var conf = srv.conf.Server.ToFiberConfig()

	conf.ErrorHandler = func(ctx fiber.Ctx, err error) error {
		// 404
		{
			if err.Error() == fmt.Sprintf("Cannot %s %s", ctx.Method(), ctx.Path()) {
				if err = http_rest_api_io.WriteError(ctx, common_errors.RouteNotFound_RestAPI()); err != nil {
					srv.components.Logger.Error().
						Format("The error response could not be recorded: '%s'. ", err).Write()

					return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
				}

				return nil
			}
		}

		// Internal server
		{
			if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
				srv.components.Logger.Error().
					Format("The error response could not be recorded: '%s'. ", err).Write()

				return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
			}

			return nil
		}
	}

	srv.app = fiber.New(conf)

	// Маршрутизатор
	{
		var location = path.Join(srv.conf.Server.Location, srv.conf.Server.Name, srv.conf.Server.Version)

		srv.router = srv.app.Group(location)
	}

	// Промежуточные слои
	{
		if srv.conf.Middlewares != nil &&
			srv.conf.Middlewares.Cors != nil &&
			srv.conf.Middlewares.Cors.Enable {
			srv.app.Use(cors_middleware.New(srv.conf.Middlewares.Cors.ToFiberConfig()))
		}
	}

	if err = srv.registerBaseRoutes(); err != nil {
		return
	}

	if err = srv.registerRoutes(); err != nil {
		return err
	}

	if err = srv.app.ShutdownWithContext(srv.ctx); err != nil {
		return
	}

	srv.components.Logger.Info().
		Text("The fiber http server has been initialized. ").Write()

	return
}
