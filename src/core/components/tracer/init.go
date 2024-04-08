package tracer

import "sm-box/src/core/components/tracer/logger"

func init() {
	config = new(Config).Default()

	if _, err := logger.New(config.Logger); err != nil {
		panic(err)
	}
}
