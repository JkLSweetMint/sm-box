package graphql

import "context"

type Engine interface {
	Serve() (err error)
	Shutdown() (err error)
}

func New(ctx context.Context) (eng Engine, err error) {
	var e = new(engine)

	eng = e

	return
}
