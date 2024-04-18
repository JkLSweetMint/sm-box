package graphql

import (
	"github.com/gofiber/fiber/v3"
)

func (e *engine) initRouter() (err error) {
	var router = e.app.Group("/dashboard")

	router.Get("/", func(ctx fiber.Ctx) error {

		return nil
	})

	return
}
