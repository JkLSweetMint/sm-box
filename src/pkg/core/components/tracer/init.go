package tracer

import (
	"sm-box/src/pkg/core/components/configurator"
	"sm-box/src/pkg/core/components/tracer/logger"
)

// Init - инициализация компонента ведения журнала трессировки вызовов функций/методов.
func Init() (err error) {
	var (
		c    configurator.Configurator[*Config]
		conf = new(Config)
	)

	if c, err = configurator.New[*Config](conf); err != nil {
		return
	} else if err = c.Private().Profile(confProfile).Read(); err != nil {
		return
	}

	if _, err = logger.New(conf.Logger); err != nil {
		return
	}

	config = conf

	return
}
