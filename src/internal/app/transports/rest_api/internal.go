package rest_api

import (
	"context"
	"github.com/gofiber/fiber/v3"
)

type engine struct {
	app *fiber.App
	ctx context.Context
}

func (e *engine) Serve() (err error) {

	if err = e.initRouter(); err != nil {
		return
	}

	if err = e.app.ShutdownWithContext(e.ctx); err != nil {
		return
	}

	if err = e.app.Listen(":8080"); err != nil {
		return
	}

	return
}

func (e *engine) Shutdown() (err error) {

	if err = e.app.Shutdown(); err != nil {
		return
	}

	return
}
