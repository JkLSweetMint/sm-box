package rest_api

import (
	"context"
	"github.com/gofiber/fiber/v3"
)

type Engine interface {
	Serve() (err error)
	Shutdown() (err error)
}

func New(ctx context.Context) (eng Engine, err error) {
	var e = new(engine)

	e.app = fiber.New()

	eng = e

	return
}
