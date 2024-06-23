package http_proxy

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	error_list "sm-box/internal/common/errors"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	"sm-box/pkg/core/components/tracer"
)

// initFiberApp - инициализация http сервера fiber.
func (eng *engine) initFiberApp() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("Starting the initialization of the fiber http server... ").
		Field("config", eng.conf.Engine.ToFiberConfig()).Write()

	var conf = eng.conf.Engine.ToFiberConfig()

	conf.ErrorHandler = func(ctx fiber.Ctx, err error) error {
		// 404
		{
			if err.Error() == fmt.Sprintf("Cannot %s %s", ctx.Method(), ctx.Path()) {
				if err = rest_api_io.WriteError(ctx, error_list.RouteNotFound_RestAPI()); err != nil {
					eng.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return nil
			}
		}

		// Ошибка с проксированием
		{
			var re = `(dial\ tcp4\ [\s\S]+:\ connectex:\ No\ connection\ could\ be\ made\ because\ the\ target\ machine\ actively\ refused\ it\.)`

			if matched, _ := regexp.MatchString(re, err.Error()); matched {
				if err = rest_api_io.WriteError(ctx, error_list.RemoteServerIsUnavailableForProxing_RestAPI()); err != nil {
					eng.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return nil
			}
		}

		// Internal server
		{
			if err = rest_api_io.WriteError(ctx, error_list.InternalServerError_RestAPI()); err != nil {
				eng.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
			}
			return nil
		}
	}

	eng.app = fiber.New(conf)

	// Маршрутизатор
	{
		var prefix string

		if eng.conf.Engine.Name != "" {
			prefix += "/" + eng.conf.Engine.Name
		}

		if eng.conf.Engine.Version != "" {
			prefix += "/" + eng.conf.Engine.Version
		}

		eng.router = eng.app.Group(prefix)
	}

	eng.initProxyRoutes()

	if err = eng.app.ShutdownWithContext(eng.ctx); err != nil {
		return
	}

	eng.components.Logger.Info().
		Text("The fiber http server has been initialized. ").Write()

	return
}
