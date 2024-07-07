package http_rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	error_list "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/transport/http_rest_api/io"
	"time"
)

// registerBaseRoutes - регистрация базовых маршрутов сервера.
func (srv *server) registerBaseRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Starting initialization of basic http rest api server routes... ").Write()

	var router = srv.router

	// /sys
	{
		var router = router.Group("/sys")

		// GET /ping
		{
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
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return http_rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			})
		}

		// GET /health
		{
			router.Get("/health", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					SystemName string `json:"system_name" xml:"SystemName"`
					Mode       string `json:"mode"        xml:"Mode"`
					Version    string `json:"version"     xml:"Version"`
					OS         string `json:"os"          xml:"OS"`

					LaunchTime string `json:"launch_time" xml:"LaunchTime"`
					UpTime     string `json:"up_time"     xml:"UpTime"`
				}

				var (
					response = new(Response)
				)

				// Обработка данных
				{
					response.SystemName = env.Vars.SystemName
					response.Mode = env.Mode.String()
					response.Version = env.Version
					response.OS = env.OS

					response.LaunchTime = env.Vars.LaunchTime.UTC().Format(time.RFC3339Nano)
					response.UpTime = time.Now().Sub(env.Vars.LaunchTime).String()
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return http_rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			})
		}

		// GET /error
		{
			router.Get("/error", func(ctx fiber.Ctx) (err error) {
				var (
					response c_errors.RestAPI
				)

				// Обработка данных
				{
					response = c_errors.ToRestAPI(error_list.Unknown())
					response.SetError(errors.New("Test. "))
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.WriteError(ctx, response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return http_rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			})
		}
	}

	srv.components.Logger.Info().
		Text("The basic http rest api server routes are initialized. ").Write()
}
