package rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	http_io "sm-box/src/pkg/tools/http/io"
)

func (eng *engine) initBaseRoutes() {
	var router = eng.router.Group("/sys")

	router.Get("/ping", func(ctx fiber.Ctx) (err error) {
		type Response struct {
			Message string `json:"message" xml:"Message" yaml:"Message"`
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
			return http_io.Write(ctx.Status(fiber.StatusOK), response)
		}
	})

	router.Get("/health", func(ctx fiber.Ctx) (err error) {
		ctx.Status(fiber.StatusOK)
		return
	})

	router.Get("/error", func(ctx fiber.Ctx) (err error) {
		return http_io.WriteError(ctx.Status(fiber.StatusInternalServerError), errors.New("Test. "))
	})
}
