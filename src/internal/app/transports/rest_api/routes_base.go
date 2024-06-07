package rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/app/transports/rest_api/io"
	error_list "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

// initBaseRoutes - регистрация базовых маршрутов системы.
func (eng *engine) initBaseRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("Starting initialization of basic http rest api routes... ").Write()

	// /sys
	{
		var router = eng.router.Group("/sys")

		router.Get("/ping", func(ctx fiber.Ctx) (err error) {
			type Response struct {
				Message string `json:"message" xml:"Message"`
			}

			var (
				response = new(Response)
			)

			// Обработка данных
			{
				response.Message = "pong"
			}

			// Отправка ответа
			{
				if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
					eng.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
					cErr.SetError(err)

					return rest_api_io.WriteError(ctx, cErr)
				}

				return
			}
		})

		router.Get("/health", func(ctx fiber.Ctx) (err error) {
			ctx.Status(fiber.StatusOK)
			return
		})

		router.Get("/error", func(ctx fiber.Ctx) (err error) {
			var (
				response c_errors.RestAPI
			)

			// Обработка данных
			{
				response = error_list.ErrUnknown_RestAPI()
				response.SetError(errors.New("Test. "))
			}

			// Отправка ответа
			{
				if err = rest_api_io.WriteError(ctx, response); err != nil {
					eng.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
					cErr.SetError(err)

					return rest_api_io.WriteError(ctx, cErr)
				}

				return
			}
		})
	}

	eng.components.Logger.Info().
		Text("The basic http rest api routes are initialized. ").Write()
}
