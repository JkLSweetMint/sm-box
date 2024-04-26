package tracer

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer/logger"
)

// Init - инициализация компонента ведения журнала трессировки вызовов функций/методов.
func Init() (err error) {
	var (
		c    configurator.Configurator[*Config]
		conf = new(Config).Default()
	)

	if c, err = configurator.New[*Config](conf); err != nil {
		return
	} else if err = c.Private().Profile(confProfile).Init(); err != nil {
		return
	}

	if _, err = logger.New(conf.Logger); err != nil {
		return
	}

	config = conf

	return
}
